package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/sean0427/micro-service-pratice-user-domain"
	mock "github.com/sean0427/micro-service-pratice-user-domain/mock/postgressql"
)

var testAuthService_AuthenticateCases = []struct {
	name string
	arg  struct {
		name     string
		password string
	}
	returnedExam bool
	returnErr    error
	wantErr      bool
}{
	{
		name: "happy",
		arg: struct {
			name     string
			password string
		}{
			name:     "any",
			password: "any2",
		},
		returnedExam: true,
		returnErr:    nil,
		wantErr:      false,
	},
	{
		name: "err",
		arg: struct {
			name     string
			password string
		}{
			name:     "any",
			password: "any2",
		},
		returnedExam: false,
		returnErr:    errors.New("any error"),
		wantErr:      true,
	},
}

func TestAuthService_Authenticate(t *testing.T) {
	for _, c := range testAuthService_AuthenticateCases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockauthRepo(ctrl)

			s := NewAuthService(m)

			m.EXPECT().
				ExamUserPassword(gomock.Any(), c.arg.name, c.arg.name).
				Return(c.returnedExam, c.returnErr).
				Times(1)

			result, err := s.Authenticate(context.Background(), c.arg.name, c.arg.password)

			if c.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if result != c.returnedExam {
				t.Errorf("expected %t, got %t", c.returnedExam, result)
			}
		})
	}
}
