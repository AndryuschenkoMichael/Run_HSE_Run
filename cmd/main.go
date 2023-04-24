package main

import (
	runHse "Run_Hse_Run"
	"Run_Hse_Run/pkg/handler"
	"Run_Hse_Run/pkg/mailer"
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

	authRepo := repository.NewAuthPostgres(db)
	gameRepo := repository.NewGamePostgres(db)
	userRepo := repository.NewUsersPostgres(db)
	friendsRepo := repository.NewFriendPostgres(db)

	mailers := mailer.NewEmailSender(dialer)
	websockets := websocket.NewGorillaServer()

	authSvc := service.NewAuthService(authRepo)
	userSvc := service.NewUsersService(userRepo)
	gameSvc := service.NewGameService(gameRepo, websockets, userSvc)
	senderSvc := service.NewSenderService(mailers)
	friendsSvc := service.NewFriendsService(friendsRepo)

	h := handler.NewHandler(authSvc, friendsSvc, gameSvc, senderSvc, userSvc)

	srv := new(runHse.Server)
	if err := srv.Run(viper.GetString("port"), h.NewMuxRoutes()); err != nil {
		log.Fatalf("Error in running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
