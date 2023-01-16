package net

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type service interface {
	Get(context.Context, *api_model.GetUsersParams) ([]*model.User, error)
	GetByID(context.Context, int64) (*model.User, error)
	Create(context.Context, *api_model.CreateUserParams) (int64, error)
	Update(context.Context, int64, *api_model.UpdateUserParams) (*model.User, error)
	Delete(context.Context, int64) error
}

type handler struct {
	service service
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := api_model.GetUsersParams{
		Name: api_model.StringToPointer(r.URL.Query().Get("name")),
	}

	users, err := h.service.Get(r.Context(), &params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
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

	id, err := getNuberId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pramas := api_model.CreateUserParams{}
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

	id, err := getNuberId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pramas := api_model.UpdateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&pramas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = h.service.Update(r.Context(), id, &pramas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := getNuberId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
