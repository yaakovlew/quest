package handler

import (
	"net/http"
	"quest/pkg/service"
)

type Quest interface {
	CreateQuest(w http.ResponseWriter, r *http.Request)
	UpdateQuest(w http.ResponseWriter, r *http.Request)
	DeleteQuest(w http.ResponseWriter, r *http.Request)
	GetQuest(w http.ResponseWriter, r *http.Request)
	GetListQuest(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Quest
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Quest: NewQuestHandler(service),
	}
}
