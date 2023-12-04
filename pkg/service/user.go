package service

import (
	"quest/pkg/models"
	"quest/pkg/repo"
)

type UserService struct {
	repo repo.User
}

func NewUserService(repo repo.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user models.User) (int, error) {
	repoUser := s.userToRepoUser(user)
	return s.repo.CreateUser(repoUser)
}

func (s *UserService) FindUser(tgUserId string) (models.User, error) {
	repoUser, err := s.repo.FindUser(tgUserId)
	if err != nil {
		return models.User{}, err
	}

	return s.repoUserToUser(repoUser), nil
}

func (s *UserService) repoUserToUser(user models.RepoUser) models.User {
	return models.User{
		TgUserId: user.TgUserId,
		Name:     user.Name,
		Age:      user.Age,
		Phone:    user.Phone,
	}
}

func (s *UserService) userToRepoUser(user models.User) models.RepoUser {
	return models.RepoUser{
		TgUserId: user.TgUserId,
		Name:     user.Name,
		Age:      user.Age,
		Phone:    user.Phone,
	}
}
