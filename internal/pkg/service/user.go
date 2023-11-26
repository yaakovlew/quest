package service

import (
	"tg-bot/internal/pkg/repo"
)

type UserService struct {
	Repo *repo.Repository
}

func NewUserService(repo *repo.Repository) *UserService {
	return &UserService{
		Repo: repo,
	}
}
