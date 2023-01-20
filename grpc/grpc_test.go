package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	"github.com/sean0427/micro-service-pratice-user-domain/grpc"
	pb "github.com/sean0427/micro-service-pratice-user-domain/grpc/grpc"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

var testListUserCasese = []struct {
	name          string
	request       *pb.UserRequest
	expectTime    int
	returnedUsers []*model.User
	returnedErr   error
	wantErr       bool
}{
	{
		name: "happy",
		request: &pb.UserRequest{
			Name: api_model.StringToPointer("any"),
		},
		expectTime: 1,
		returnedUsers: []*model.User{
			{
				ID:   1,
				Name: "any",
			},
			{
				ID:   2,
				Name: "any2",
			},
		},
		returnedErr: nil,
		wantErr:     false,
	},
	{
		name: "errpr - not found",
		request: &pb.UserRequest{
			Name: api_model.StringToPointer("any"),
		},
		expectTime:    1,
		returnedUsers: []*model.User{},
		returnedErr:   nil,
		wantErr:       true,
	}, {
		name:          "error - no name",
		request:       &pb.UserRequest{},
		expectTime:    0,
		returnedUsers: nil,
		returnedErr:   nil,
		wantErr:       true,
	},
	{
		name: "error - servce return err",
		request: &pb.UserRequest{
			Name: api_model.StringToPointer("any"),
		},
		expectTime:    1,
		returnedUsers: nil,
		returnedErr:   errors.New("any"),
		wantErr:       true,
	},
}

func TestGrpcService_ListUsers(t *testing.T) {
	for _, c := range testListUserCasese {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			m := NewMockservice(ctrl)
			m.EXPECT().
				Get(gomock.Any(), gomock.Eq(&api_model.GetUsersParams{Name: c.request.Name})).
				Return(c.returnedUsers, c.returnedErr).
				Times(c.expectTime)

			grpc := grpc.New(m)

			res, err := grpc.ListUsers(context.Background(), c.request)

			if c.wantErr && err != nil {
				if c.returnedErr != nil && err == nil {
					t.Fatalf("expect err %s", err.Error())
				}
				return
			}

			if v, e := len(res.Users), len(c.returnedUsers); v != e {
				t.Fatalf("ecpect %d but %d", v, e)
			}
		})
	}
}

var testGetUserCase = []struct {
	name         string
	request      *pb.UserRequest
	expectTime   int
	returnedUser *model.User
	returnedErr  error
	wantErr      bool
}{
	{
		name: "happy",
		request: &pb.UserRequest{
			Id: api_model.Int64ToPointer(11),
		},
		expectTime: 1,
		returnedUser: &model.User{
			ID:   1,
			Name: "any",
		},
		returnedErr: nil,
		wantErr:     false,
	},
	{
		name: "error - not found",
		request: &pb.UserRequest{
			Id: api_model.Int64ToPointer(2211),
		},
		expectTime:   1,
		returnedUser: nil,
		returnedErr:  nil,
		wantErr:      true,
	}, {
		name:         "error - no id",
		request:      &pb.UserRequest{},
		expectTime:   0,
		returnedUser: nil,
		returnedErr:  nil,
		wantErr:      true,
	},
	{
		name: "error - servce return err",
		request: &pb.UserRequest{
			Id: api_model.Int64ToPointer(11),
		},
		expectTime:   1,
		returnedUser: nil,
		returnedErr:  errors.New("any"),
		wantErr:      true,
	},
}

func TestGrpcService_GetUser(t *testing.T) {
	for _, c := range testGetUserCase {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			idMatcher := gomock.Any()
			if c.request != nil && c.request.Id != nil {
				idMatcher = gomock.Eq(*c.request.Id)
			}

			m := NewMockservice(ctrl)
			m.EXPECT().
				GetByID(gomock.Any(), idMatcher).
				Return(c.returnedUser, c.returnedErr).
				Times(c.expectTime)

			grpc := grpc.New(m)

			res, err := grpc.GetUser(context.Background(), c.request)

			if c.wantErr && err != nil {
				if c.returnedErr != nil && err == nil {
					t.Fatalf("expect err %s", err.Error())
				}
				return
			}

			if res == nil {
				t.Fatal("expect not nil")
			}

			if v, e := res.User.Id, c.returnedUser.ID; v != e {
				t.Fatalf("ecpect %d but %d", v, e)
			}
		})
	}
}
