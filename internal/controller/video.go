package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/usecases"
)

type VideoController struct {
	getVideosUseCase usecases.GetVideosUseCase
	streamVideoUseCase usecases.StreamVideoUseCase
}

func NewVideoControler(getVideosUseCase usecases.GetVideosUseCase, streamVideoUseCase usecases.StreamVideoUseCase) VideoController {
	return VideoController{getVideosUseCase: getVideosUseCase, streamVideoUseCase: streamVideoUseCase}
}

func (c VideoController) GetVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := c.getVideosUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(videos)
}

func (c VideoController) StreamVideo(w http.ResponseWriter, r *http.Request) {
	videoName := strings.TrimPrefix(r.URL.Path, "/stream/")
	path := c.streamVideoUseCase.Execute(videoName)
	path = strings.ReplaceAll(path, "\\", "/")

	http.Redirect(w, r, "/"+path, http.StatusFound)
}
