package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	apptodo "github.com/xrodazxx/App-ToDO"

	"github.com/xrodazxx/App-ToDO/pkg/handler"
	"github.com/xrodazxx/App-ToDO/pkg/repository"
	"github.com/xrodazxx/App-ToDO/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs:%s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variavles:%s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db_host"),
		Port:     viper.GetString("db_port"),
		Username: viper.GetString("db_username"),
		DBName:   viper.GetString("db_dbname"),
		SSLMode:  viper.GetString("db_sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(apptodo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server:%s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
