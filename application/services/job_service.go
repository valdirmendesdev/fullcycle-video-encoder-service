package services

import (
	"errors"
	"github.com/valdirmendesdev/encoder-service/application/repositories"
	"github.com/valdirmendesdev/encoder-service/domain"
)

type JobService struct {
	Job *domain.Job
	JobRepository repositories.JobRepository
	VideoService VideoService
}

func (j *JobService) Start() error {

	err := j.changeJobStatus(JOB_STATUS_DOWNLOADING)

	if err != nil {
		return j.failJob(err)
	}

	err = j.VideoService.Download(GetInputBucketName())

	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(JOB_STATUS_FRAGMENTING)

	if err != nil {
		return j.failJob(err)
	}

	err = j.VideoService.Fragment()

	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(JOB_STATUS_ENCODING)

	if err != nil {
		return j.failJob(err)
	}

	err = j.VideoService.Encode()

	if err != nil {
		return j.failJob(err)
	}

	err = j.performUpload()

	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(JOB_STATUS_FINISHING)

	if err != nil {
		return j.failJob(err)
	}

	err = j.VideoService.Finish()

	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(JOB_STATUS_COMPLETED)

	if err != nil {
		return j.failJob(err)
	}

	return nil

}

func (j *JobService) performUpload() error {

	err := j.changeJobStatus(JOB_STATUS_UPLOADING)

	if err != nil {
		return j.failJob(err)
	}

	videoUpload := NewVideoUpload()
	videoUpload.OutputBucket = GetOutputBucketName()
	videoUpload.VideoPath = MountVideoFolderName(j.VideoService.Video.ID)
	concurrency, _ := GetNumberOfConcurrencyUploadProcesses()
	doneUpload := make(chan string)

	go videoUpload.ProcessUpload(concurrency, doneUpload)

	uploadResult := <- doneUpload

	if uploadResult != UPLOAD_COMPLETED {
		return j.failJob(errors.New(uploadResult))
	}
	
	return err
}

func (j *JobService) changeJobStatus(status string) error {
	var err error

	j.Job.Status = status
	j.Job, err = j.JobRepository.Update(j.Job)

	if err != nil {
		return j.failJob(err)
	}
	return nil
}

func (j *JobService) failJob(error error) error {

	j.Job.Status = JOB_STATUS_FAILED
	j.Job.Error = error.Error()

	_, err := j.JobRepository.Update(j.Job)

	if err != nil {
		return err
	}

	return error
}