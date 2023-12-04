package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"quest/pkg/handler"
	"quest/pkg/models"
	"quest/pkg/repo"
	"quest/pkg/server"
	"quest/pkg/service"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env varibles: %s", err.Error())
		return
	}
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
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(handlers.InitRoutes()); err != nil {
			log.Fatalf("Problem with start server, because %s", err.Error())
			return
		}
	}()

	log.Println("backend started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("backend shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
		return
	}

}
