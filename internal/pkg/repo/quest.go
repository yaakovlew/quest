package repo

import (
	"gorm.io/gorm"
	"quest/internal/models"
)

const limit = 3

type QuestRepo struct {
	db *gorm.DB
}

func NewQuestRepo(db *gorm.DB) *QuestRepo {
	return &QuestRepo{
		db: db,
	}
}

func (r *QuestRepo) CreateQuest(quest models.RepoQuest) (int, error) {
	if err := r.db.Create(quest).Error; err != nil {
		return 0, err
	}

	return int(quest.ID), nil
}

func (r *QuestRepo) UpdateQuest(quest models.RepoQuest) (models.RepoQuest, error) {
	if err := r.db.Model(&models.RepoQuest{}).Where("id=?", quest.ID).Updates(&quest).Error; err != nil {
		return models.RepoQuest{}, err
	}

	return quest, nil
}

func (r *QuestRepo) DeleteQuest(quest models.RepoQuest) (int, error) {
	if err := r.db.Delete(&quest).Error; err != nil {
		return 0, err
	}

	return int(quest.ID), nil
}

func (r *QuestRepo) GetQuest(questId int) (models.RepoQuest, error) {
	var quest models.RepoQuest
	if err := r.db.Where("id=?", questId).First(&quest).Error; err != nil {
		return models.RepoQuest{}, err
	}

	return quest, nil
}

func (r *QuestRepo) GetQuestsByPage(page int) ([]models.RepoQuest, error) {
	offset := limit * (page - 1)
	var quests []models.RepoQuest
	if err := r.db.Offset(offset).Limit(limit).Order("id").Find(&quests).Error; err != nil {
		return nil, err
	}

	return quests, nil
}
