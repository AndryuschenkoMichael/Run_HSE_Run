package main

import (
	runHse "Run_Hse_Run"
	"Run_Hse_Run/pkg/handler"
	"Run_Hse_Run/pkg/repository"
	"Run_Hse_Run/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error with initializing config file: %s", err.Error())
	}

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(runHse.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error in running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
