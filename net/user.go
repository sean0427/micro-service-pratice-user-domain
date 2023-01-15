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
	Create(context.Context, *model.CreateUserParams) (string, error)
	Update(context.Context, string, *model.UpdateUserParams) (*model.User, error)
	Delete(context.Context, string) error
}

type handler struct {
	service service
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := model.GetUsersParams{
		Name: model.StringToPointer(r.URL.Query().Get("name")),
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

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pramas := model.CreateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&pramas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := h.service.Create(r.Context(), &pramas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	pramas := model.UpdateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&pramas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := h.service.Update(r.Context(), id, &pramas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
