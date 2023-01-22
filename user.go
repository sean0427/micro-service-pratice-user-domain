package user

import (
	"context"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

type repository interface {
	Get(ctx context.Context, params *api_model.GetUsersParams) ([]*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *api_model.CreateUserParams) (int64, error)
	Update(ctx context.Context, id int64, user *api_model.UpdateUserParams) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserService struct {
	repo repository
}

func New(repo repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Get(ctx context.Context, params *api_model.GetUsersParams) ([]*model.User, error) {
	return s.repo.Get(ctx, params)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, params *api_model.CreateUserParams) (int64, error) {
	return s.repo.Create(ctx, params)
}

func (s *UserService) Update(ctx context.Context, id int64, params *api_model.UpdateUserParams) (*model.User, error) {
	return s.repo.Update(ctx, id, params)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
