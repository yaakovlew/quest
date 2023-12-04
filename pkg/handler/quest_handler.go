package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"quest/pkg/models"
	"quest/pkg/service"
	"strconv"
)

type QuestHandler struct {
	service *service.Service
}

func NewQuestHandler(service *service.Service) *QuestHandler {
	return &QuestHandler{
		service: service,
	}
}

func (h *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest models.Quest
	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	questId, err := h.service.CreateQuest(quest)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(map[string]interface{}{
		"quest_id": questId,
	})
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *QuestHandler) UpdateQuest(w http.ResponseWriter, r *http.Request) {
	questStr := mux.Vars(r)["id"]

	questId, err := strconv.Atoi(questStr)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	var quest models.Quest
	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	updatedQuest, err := h.service.UpdateQuest(questId, quest)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(updatedQuest)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *QuestHandler) DeleteQuest(w http.ResponseWriter, r *http.Request) {
	quest := mux.Vars(r)["id"]

	questId, err := strconv.Atoi(quest)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	deletedQuestId, err := h.service.DeleteQuest(questId)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(map[string]interface{}{
		"quest_id": deletedQuestId,
	})
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *QuestHandler) GetQuest(w http.ResponseWriter, r *http.Request) {
	questStr := mux.Vars(r)["id"]

	questId, err := strconv.Atoi(questStr)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	quest, err := h.service.GetQuest(questId)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(quest)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *QuestHandler) GetListQuest(w http.ResponseWriter, r *http.Request) {
	page := mux.Vars(r)["page"]

	pageId, err := strconv.Atoi(page)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	quests, err := h.service.GetQuestsByPage(pageId)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(quests)
	if err != nil {
		httpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
