package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"quest/pkg/models"
	"quest/pkg/msgHandler"
	"quest/pkg/repo"
	"quest/pkg/service"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env varibles: %s", err.Error())
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatalf("error upload tg token: %s", err.Error())
		return
	}
	bot.Debug = true

	db, err := repo.NewPostgresDB(repo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("Fatal to connect to DB, because: %s", err.Error())
		return
	}

	db.AutoMigrate(&models.RepoQuest{}, &models.RepoUser{})

	repos := repo.NewRepository(db)
	services := service.NewService(repos)
	handler := msgHandler.NewMSGHandler(services, bot)

	log.Printf("bot started %s", bot.Self.UserName)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go handler.ValidateMessage()

	<-signalChannel
	log.Printf("bot closed %s", bot.Self.UserName)
}
