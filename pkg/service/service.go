package service

import (
	"quest/pkg/models"
	"quest/pkg/repo"
)

type User interface {
	CreateUser(user models.User) (int, error)
	FindUser(tgUserId string) (models.User, error)
	repoUserToUser(user models.RepoUser) models.User
	userToRepoUser(user models.User) models.RepoUser
}

type Quest interface {
	CreateQuest(quest models.Quest) (int, error)
	UpdateQuest(id int, quest models.Quest) (models.Quest, error)
	DeleteQuest(id int) (int, error)
	GetQuest(questId int) (models.Quest, error)
	GetQuestsByPage(page int) ([]models.Quest, error)
	repoQuestToQuest(models.RepoQuest) models.Quest
	questToRepoQuest(models.Quest) models.RepoQuest
}

type Service struct {
	User
	Quest
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		User:  NewUserService(repo.User),
		Quest: NewQuestService(repo.Quest),
	}
}
