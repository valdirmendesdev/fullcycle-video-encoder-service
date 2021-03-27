package services

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/valdirmendesdev/encoder-service/application/repositories"
	"github.com/valdirmendesdev/encoder-service/domain"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

const tmpDirectoryPath = "localStoragePath"

func mountTempFilename(filename, extension string ) string {
	return os.Getenv(tmpDirectoryPath) + "/" + filename + extension
}

func mountTempFolder(foldername string) string {
	return os.Getenv(tmpDirectoryPath) + "/" + foldername
}

func (v *VideoService) Download(bucketName string) error {
	ctx :=  context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)

	if err != nil {
		return err
	}

	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(mountTempFilename(v.Video.ID, ".mp4"))
	if err != nil {
		return nil
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	defer f.Close()

	log.Printf("video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {
	err := os.Mkdir(mountTempFolder(v.Video.ID), os.ModePerm)
	if err != nil {
		return err
	}

	source := mountTempFilename(v.Video.ID, ".mp4")
	target := mountTempFilename(v.Video.ID, ".frag")

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	printOutput(output)
	return nil
}

func (v *VideoService) Encode() error {

	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, mountTempFilename(v.Video.ID, ".frag"))
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, mountTempFolder(v.Video.ID))
	cmdArgs = append(cmdArgs, "-f")
	cmdArgs = append(cmdArgs, "--exec-dir")
	cmdArgs = append(cmdArgs, "/opt/bento4/bin/")
	cmd := exec.Command("mp4dash", cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	printOutput(output)
	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("=====> Output: %s\n", string(out))
	}
}