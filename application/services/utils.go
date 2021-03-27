package services

import "os"

const tmpDirPath = "localStoragePath"

func MountTempFilename(filename, extension string) string {
	return GetTmpDir() + filename + extension
}

func MountTempFolder(foldername string) string {
	return GetTmpDir() + foldername
}

func GetTmpDir() string {
	return os.Getenv(tmpDirPath) + "/"
}