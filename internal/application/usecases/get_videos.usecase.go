package usecases

import "github.com/Gabriel-Schiestl/dash-streaming/internal/application/services"

type GetVideosUseCaseResponse struct {
	Videos []string `json:"videos"`
}

type GetVideosUseCase struct {
	videoService services.IVideoService
}

func NewGetVideosUseCase(videoService services.IVideoService) GetVideosUseCase {
	return GetVideosUseCase{videoService: videoService}
}

func (uc GetVideosUseCase) Execute() (*GetVideosUseCaseResponse, error) {
	videos, err := uc.videoService.GetVideos()
	if err != nil {
		return nil, err
	}

	videoJson := &GetVideosUseCaseResponse{Videos: videos}
	return videoJson, nil
}