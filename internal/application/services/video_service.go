package services

type IVideoService interface {
	GetVideos() ([]string, error)
	VerifyIfDashExists(dashPath string) bool
	CreateDash(videoPath string, dashDir string)
}