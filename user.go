package user

import (
	"context"

	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type repository interface {
	Get(ctx context.Context, params *model.GetUsersParams) ([]*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
}

type UserService struct {
	repo repository
}

func New(repo repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Get(ctx context.Context, params *model.GetUsersParams) ([]*model.User, error) {
	return s.repo.Get(ctx, params)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}
