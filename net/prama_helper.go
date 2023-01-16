package net

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func getNuberId(r *http.Request) (int64, error) {
	id := chi.URLParam(r, "id")
	_id, err := strconv.Atoi(id)

	return int64(_id), err
}
