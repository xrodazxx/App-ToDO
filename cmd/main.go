package main

import (
	"log"

	apptodo "github.com/xrodazxx/App-ToDO"

	"github.com/xrodazxx/App-ToDO/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(apptodo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server:%s", err.Error())
	}

}
