package domain_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/encoder-service/domain"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	require.Error(t, video.Validate())
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	require.Error(t, video.Validate())
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "resourceId"
	video.FilePath = "/file"
	video.CreatedAt = time.Now()
	require.Nil(t, video.Validate())
}