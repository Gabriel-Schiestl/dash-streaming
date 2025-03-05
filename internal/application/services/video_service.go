package services

type IVideoService interface {
	GetVideos() ([]string, error)
}