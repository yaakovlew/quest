package handler

import "github.com/go-chi/chi"

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/quest", func(r chi.Router) {
		r.Post("", h.CreateQuest)
		r.Delete("/{id}", h.DeleteQuest)
		r.Put("", h.UpdateQuest)
		r.Get("/{id}", h.GetQuest)
		r.Get("/page/{page}", h.GetListQuest)
	})

	return r
}
