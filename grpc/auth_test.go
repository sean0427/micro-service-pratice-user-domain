package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/sean0427/micro-service-pratice-user-domain/grpc"
	pb "github.com/sean0427/micro-service-pratice-user-domain/grpc/auth"
	mock "github.com/sean0427/micro-service-pratice-user-domain/mock"
)

var testAuthGrpc_AuthenticateCases = []struct {
	name            string
	req             *pb.AuthRequest
	wantErr         bool
	retrunedSuccess bool
	retrunedError   error
	times           int
}{
	{
		name: "happy",
		req: &pb.AuthRequest{
			Name:     "any1",
			Password: "any2",
		},
		retrunedSuccess: true,
		retrunedError:   nil,
		times:           1,
	},
	{
		name: "failed",
		req: &pb.AuthRequest{
			Name:     "any1",
			Password: "any2",
		},
		retrunedSuccess: false,
		retrunedError:   nil,
		wantErr:         false,
		times:           1,
	},
	{
		name: "error - no name",
		req: &pb.AuthRequest{
			Password: "any2",
		},
		retrunedSuccess: false,
		retrunedError:   nil,
		wantErr:         true,
		times:           0,
	},
	{
		name: "error - no pass",
		req: &pb.AuthRequest{
			Name: "any1",
		},
		retrunedSuccess: false,
		retrunedError:   nil,
		wantErr:         true,
		times:           0,
	},
	{
		name:            "failed - nil req",
		req:             nil,
		retrunedSuccess: false,
		retrunedError:   nil,
		wantErr:         true,
		times:           0,
	},
	{
		name: "error",
		req: &pb.AuthRequest{
			Name:     "any1",
			Password: "any2",
		},
		retrunedSuccess: false,
		retrunedError:   errors.New("any"),
		wantErr:         true,
		times:           1,
	},
}

func TestAuthGrpc_Authenticate(t *testing.T) {
	for _, tt := range testAuthGrpc_AuthenticateCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := mock.NewMockauthSerivice(ctrl)

			nameMacher := gomock.Any()
			passwordMacher := gomock.Any()

			if tt.req != nil {
				nameMacher = gomock.Eq(tt.req.Name)
				passwordMacher = gomock.Eq(tt.req.Password)
			}

			m.EXPECT().
				Authenticate(gomock.Any(), nameMacher, passwordMacher).
				Return(tt.retrunedSuccess, tt.retrunedError).
				Times(tt.times)

			handler := NewAuthHandler(m)
			got, err := handler.Authenticate(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AuthGrpc.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got.Success != tt.retrunedSuccess {
				t.Errorf("AuthGrpc.Authenticate() = %v, want %v", got, tt.retrunedSuccess)
			}
		})
	}
}
