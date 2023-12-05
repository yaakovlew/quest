package handler

import (
	"github.com/gorilla/mux"
)

func (h *Handler) InitRoutes() *mux.Router {

	r := mux.NewRouter()
	questRouter := r.PathPrefix("/quest").Subrouter()
	questRouter.Use(checkAuthHeader)
	questRouter.HandleFunc("", h.CreateQuest).Methods("POST")
	questRouter.HandleFunc("/{id}", h.DeleteQuest).Methods("DELETE")
	questRouter.HandleFunc("/{id}", h.UpdateQuest).Methods("PUT")
	questRouter.HandleFunc("/{id}", h.GetQuest).Methods("GET")
	questRouter.HandleFunc("/page/{page}", h.GetListQuest).Methods("GET")

	return r
}
