package repositories_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/encoder-service/application/repositories"
	"github.com/valdirmendesdev/encoder-service/domain"
	"github.com/valdirmendesdev/encoder-service/framework/database"
	"testing"
	"time"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID,video.ID)
}
