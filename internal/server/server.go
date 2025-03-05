package server

import (
	"fmt"
	"net/http"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/usecases"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/controller"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/infra/services"
)

func Init() {
	service := services.NewVideoService()
	useCase := usecases.NewGetVideosUseCase(service)
	videoController := controller.NewVideoControler(useCase)

	http.HandleFunc("/videos", videoController.GetVideos)

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}