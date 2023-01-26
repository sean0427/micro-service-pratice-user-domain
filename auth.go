package user

import "context"

type authRepo interface {
	ExamUserPassword(ctx context.Context, name, password string) (bool, error)
}

type AuthService struct {
	repo authRepo
}

func NewAuthService(repo authRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, name, password string) (bool, error) {
	return true, nil
}
