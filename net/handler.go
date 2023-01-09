package net

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type service interface {
	Get(context.Context, *model.GetUsersParams) ([]*model.User, error)
	GetByID(context.Context, string) (*model.User, error)
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := model.GetUsersParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := h.service.Get(r.Context(), &params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) InitHandler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Get("/:id", h.GetById)
	})

	return r
}
