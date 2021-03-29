package services_test

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/encoder-service/application/services"
	"log"
	"testing"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestVideoServiceUpload(t *testing.T) {

	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download(bucketNameToTest)
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = bucketNameToTest
	videoUpload.VideoPath = services.MountVideoFolderName(video.ID)

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(50, doneUpload)

	result := <-doneUpload
	require.Equal(t, result, services.UPLOAD_COMPLETED)

	err = videoService.Finish()
	require.Nil(t, err)
}
