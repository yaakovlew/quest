package repo

import (
	"gorm.io/gorm"
	"quest/pkg/models"
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
	if err := r.db.Create(&quest).Error; err != nil {
		return 0, err
	}

	return int(quest.ID), nil
}

func (r *QuestRepo) UpdateQuest(id int, quest models.RepoQuest) (models.RepoQuest, error) {
	if err := r.db.Model(&models.RepoQuest{}).Where("id=?", id).Updates(&quest).Error; err != nil {
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

func (r *QuestRepo) GetPageAmount() (int, error) {
	var amountOfQuests int64
	if err := r.db.Model(&models.RepoQuest{}).Count(&amountOfQuests).Error; err != nil {
		return 0, err
	}

	pages := int(amountOfQuests) / limit

	if amountOfQuests%limit != 0 {
		pages++
	}

	return pages, nil
}
