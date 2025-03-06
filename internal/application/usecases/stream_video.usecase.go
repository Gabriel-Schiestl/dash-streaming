package usecases

import (
	"os"
	"path/filepath"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/services"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/constants"
)

type StreamVideoUseCaseResponse struct {
	Videos []string `json:"videos"`
}

type StreamVideoUseCase struct {
	videoService services.IVideoService
}

func NewStreamVideosUseCase(videoService services.IVideoService) StreamVideoUseCase {
	return StreamVideoUseCase{videoService: videoService}
}

func (uc StreamVideoUseCase) Execute(videoName string) string {
	videoPath := filepath.Join(constants.VideosDir, videoName)
	dashDir := filepath.Join(constants.DashDir, videoName)

	if uc.videoService.VerifyIfDashExists(dashDir) {
		return filepath.Join("dash", videoName, "manifest.mpd")
	}

	os.Mkdir(dashDir, os.ModePerm)

	uc.videoService.CreateDash(videoPath, dashDir)

	return filepath.Join("dash", videoName, "manifest.mpd")
}