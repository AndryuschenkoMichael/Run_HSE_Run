package main

import (
	runHse "Run_Hse_Run"
	"Run_Hse_Run/pkg/emailer"
	"Run_Hse_Run/pkg/handler"
	"Run_Hse_Run/pkg/repository"
	"Run_Hse_Run/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error with initializing config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error with initializing environment file: %s", err.Error())
	}

	dialer := gomail.NewDialer(
		viper.GetString("emailer.host"),
		viper.GetInt("emailer.port"),
		viper.GetString("emailer.email"),
		os.Getenv("EMAIL_PASSWORD"))

	emailers := emailer.NewEmailSender(dialer)
	repositories := repository.NewRepository()
	services := service.NewService(repositories, emailers)
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
