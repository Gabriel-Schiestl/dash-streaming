package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Gabriel-Schiestl/dash-streaming/internal/application/usecases"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/constants"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/controller"
	"github.com/Gabriel-Schiestl/dash-streaming/internal/infra/services"
)

func Init() {
	service := services.NewVideoService()
	getVideosUseCase := usecases.NewGetVideosUseCase(service)
	streamVideoUseCase := usecases.NewStreamVideosUseCase(service)
	videoController := controller.NewVideoControler(getVideosUseCase, streamVideoUseCase)

	os.Mkdir(constants.DashDir, os.ModePerm)

	http.HandleFunc("/videos", videoController.GetVideos)
	http.HandleFunc("/stream/", videoController.StreamVideo)

	fs := http.FileServer(http.Dir(constants.DashDir))
	http.Handle("/dash/", http.StripPrefix("/dash/", fs))

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}