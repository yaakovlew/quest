package repo

import (
	"gorm.io/gorm"
	"quest/pkg/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user models.RepoUser) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return int(user.ID), nil
}

func (r *UserRepo) FindUser(tgUserId int) (models.RepoUser, error) {
	var user models.RepoUser
	if err := r.db.Where("tg_user_id=?", tgUserId).First(&user).Error; err != nil {
		return models.RepoUser{}, err
	}

	return user, nil
}
