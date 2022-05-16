package main

import (
	runHse "Run_Hse_Run"
	"Run_Hse_Run/pkg/handler"
	"Run_Hse_Run/pkg/mailer"
	"Run_Hse_Run/pkg/queue"
	"Run_Hse_Run/pkg/repository"
	"Run_Hse_Run/pkg/service"
	"Run_Hse_Run/pkg/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
		viper.GetString("mailer.host"),
		viper.GetInt("mailer.port"),
		viper.GetString("mailer.email"),
		os.Getenv("MAIL_PASSWORD"))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Faild to initialize db: %s", err.Error())
	}

	mailers := mailer.NewMailer(dialer)
	repositories := repository.NewRepository(db)
	queues := queue.NewQueue()
	websockets := websocket.NewServer()
	services := service.NewService(repositories, mailers, queues, websockets)
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
