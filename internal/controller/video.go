package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/usecases"
)

type VideoController struct {
	usecase usecases.GetVideosUseCase
}

func NewVideoControler(usecase usecases.GetVideosUseCase) VideoController {
	return VideoController{usecase: usecase}
}

func (c VideoController) GetVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := c.usecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(videos)
}
