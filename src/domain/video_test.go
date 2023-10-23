package domain_test

import (
	"testing"
	"time"

	"github.com/elisiomualumene/video_encoder/src/domain"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmpty(t *testing.T){
	video := domain.NewVideo()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T){
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIsValid(t *testing.T){
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}