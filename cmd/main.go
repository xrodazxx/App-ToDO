package main

import (
	"log"

	apptodo "github.com/xrodazxx/App-ToDO"

	"github.com/xrodazxx/App-ToDO/pkg/handler"
	"github.com/xrodazxx/App-ToDO/pkg/repository"
	"github.com/xrodazxx/App-ToDO/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(apptodo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server:%s", err.Error())
	}

}
