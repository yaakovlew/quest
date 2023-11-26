package repo

import (
	"gorm.io/gorm"
	"tg-bot/internal/pkg/models"
)

type User interface {
	CreateUser(user models.RepoUser) (int, error)
	FindUser(tgUserId string) (models.RepoUser, error)
}

type Quest interface {
	CreateQuest(quest models.RepoQuest) (int, error)
	UpdateQuest(quest models.RepoQuest) (models.RepoQuest, error)
	DeleteQuest(quest models.RepoQuest) (int, error)
	GetQuest(questId int) (models.RepoQuest, error)
	GetQuestsByPage(page int) ([]models.RepoQuest, error)
}

type Repository struct {
	Quest
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Quest: NewQuestRepo(db),
		User:  NewUserRepo(db),
	}
}
