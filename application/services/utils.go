package services

import (
	"os"
	"strconv"
)

const UPLOAD_COMPLETED = "upload completed"

const (
	ENV_TMP_DIR_PATH      = "localStoragePath"
	ENV_INPUT_BUCKET_NAME = "inputBucketName"
	ENV_OUTPUT_BUCKET_NAME = "outputBucketName"
	ENV_CONCURRENCY_UPLOAD = "CONCURRENCY_UPLOAD"
	ENV_CONCURRENCY_WORKERS = "CONCURRENCY_WORKERS"
)

const (
	JOB_STATUS_FAILED      string = "FAILED"
	JOB_STATUS_STARTING           = "STARTING"
	JOB_STATUS_DOWNLOADING        = "DOWNLOADING"
	JOB_STATUS_FRAGMENTING        = "FRAGMENTING"
	JOB_STATUS_ENCODING           = "ENCODING"
	JOB_STATUS_UPLOADING          = "UPLOADING"
	JOB_STATUS_FINISHING          = "FINISHING"
	JOB_STATUS_COMPLETED          = "COMPLETED"
)

func MountVideoFilename(filename, extension string) string {
	return GetTmpDir() + filename + extension
}

func MountVideoFolderName(foldername string) string {
	return GetTmpDir() + foldername
}

func GetTmpDir() string {
	return os.Getenv(ENV_TMP_DIR_PATH) + "/"
}

func GetInputBucketName() string {
	return os.Getenv(ENV_INPUT_BUCKET_NAME)
}

func GetOutputBucketName() string {
	return os.Getenv(ENV_OUTPUT_BUCKET_NAME)
}

func GetNumberOfConcurrencyUploadProcesses() (int, error) {
	return strconv.Atoi(os.Getenv(ENV_CONCURRENCY_UPLOAD))
}
func GetNumberOFConcurrencyWorkers() (int, error) {
	return strconv.Atoi(os.Getenv(ENV_CONCURRENCY_WORKERS))
}