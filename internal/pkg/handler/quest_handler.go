package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"quest/internal/pkg/service"
	"strconv"
)

type QuestHandler struct {
	service service.Service
}

func NewQuestHandler(service service.Service) *QuestHandler {
	return &QuestHandler{
		service: service,
	}
}

func (h *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	r.
}

func (h *QuestHandler) UpdateQuest(w http.ResponseWriter, r *http.Request) {}

func (h *QuestHandler) DeleteQuest(w http.ResponseWriter, r *http.Request) {
/*	quest := chi.URLParam(r, "id")

	questId, err := strconv.Atoi(quest)
	if err != nil{
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	deletedQuestId, err := h.service.DeleteQuest(questId)
	if err != nil{
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(200)
	w.Write()*/
}

func (h *QuestHandler) GetQuest(w http.ResponseWriter, r *http.Request) {}

func (h *QuestHandler) GetListQuest(w http.ResponseWriter, r *http.Request) {}
