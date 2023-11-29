package service

import (
	"quest/internal/models"
	"quest/internal/pkg/repo"
)

type QuestService struct {
	repo repo.Quest
}

func NewQuestService(repo repo.Quest) *QuestService {
	return &QuestService{
		repo: repo,
	}
}

func (s *QuestService) CreateQuest(quest models.Quest) (int, error) {
	repoQuest := s.questToRepoQuest(quest)

	return s.repo.CreateQuest(repoQuest)
}

func (s *QuestService) UpdateQuest(quest models.Quest) (models.RepoQuest, error) {
	repoQuest := s.questToRepoQuest(quest)

	return s.repo.UpdateQuest(repoQuest)
}

func (s *QuestService) DeleteQuest(id int) (int, error) {
	repoQuest, err := s.repo.GetQuest(id)
	if err != nil{
		return 0, err
	}

	return s.repo.DeleteQuest(repoQuest)
}

func (s *QuestService) GetQuest(questId int) (models.Quest, error) {
	repoQuest, err := s.repo.GetQuest(questId)
	if err != nil {
		return models.Quest{}, err
	}

	return s.repoQuestToQuest(repoQuest), err
}

func (s *QuestService) GetQuestsByPage(page int) ([]models.Quest, error) {
	repoQuests, err := s.repo.GetQuestsByPage(page)
	if err != nil {
		return nil, err
	}

	var quests []models.Quest
	for _, quest := range repoQuests {
		quests = append(quests, s.repoQuestToQuest(quest))
	}

	return quests, nil
}

func (s *QuestService) repoQuestToQuest(quest models.RepoQuest) models.Quest {
	return models.Quest{
		Name:          quest.Name,
		Description:   quest.Description,
		AuthorComment: quest.AuthorComment,
		Point:         quest.Point,
		AgeLevel:      quest.AgeLevel,
	}
}

func (s *QuestService) questToRepoQuest(quest models.Quest) models.RepoQuest {
	return models.RepoQuest{
		Name:          quest.Name,
		Description:   quest.Description,
		AuthorComment: quest.AuthorComment,
		Point:         quest.Point,
		AgeLevel:      quest.AgeLevel,
	}
}
