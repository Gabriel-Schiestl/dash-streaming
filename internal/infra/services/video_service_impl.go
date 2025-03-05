package services

import (
	"errors"
	"os"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/services"
)

type videoService struct {
}

func NewVideoService() services.IVideoService {
	return videoService{}
}

func (v videoService) GetVideos() ([]string, error) {
	videos, err := os.ReadDir("videos")
	if err != nil {
		return nil, errors.New("Error reading videos directory: " + err.Error())
	}

	var videosNames []string

	for _, video := range videos {
		videosNames = append(videosNames, video.Name())
	}

	return videosNames, nil
}