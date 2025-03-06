package services

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/services"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/constants"
)

type videoService struct {
}

func NewVideoService() services.IVideoService {
	return videoService{}
}

func (v videoService) GetVideos() ([]string, error) {
	videos, err := os.ReadDir(constants.VideosDir)
	if err != nil {
		return nil, errors.New("Error reading videos directory: " + err.Error())
	}

	var videosNames []string

	for _, video := range videos {
		videosNames = append(videosNames, video.Name())
	}

	return videosNames, nil
}

func (v videoService) VerifyIfDashExists(dashPath string) bool {
	_, err := os.Stat(dashPath)
	return err == nil
}

func (v videoService) CreateDash(videoPath string, dashDir string) {
	cmd := exec.Command("C:/Program Files/GPAC/mp4box.exe",
		"-dash", "8000",
		"-frag", "8000",
		"-segment-name", "segment_",
		"-out", filepath.Join(dashDir, "manifest.mpd"),
		videoPath,
	)

	_, err := cmd.CombinedOutput()
	if err != nil {
		panic("Error creating dash: " + err.Error())
	}
}