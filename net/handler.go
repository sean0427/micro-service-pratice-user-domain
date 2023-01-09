package net

import (
	"github.com/go-chi/chi/v5"
)

func New(service service) *handler {
	return &handler{
		service: service}
}

func (h *handler) InitHandler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Get("/:id", h.GetById)
	})

	return r
}
