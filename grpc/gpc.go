package grpc

import (
	"context"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	"github.com/sean0427/micro-service-pratice-user-domain/grpc/grpc"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(s service) *GrpcService {
	return &GrpcService{
		service: s,
	}
}

type GrpcService struct {
	grpc.UnimplementedUserHandlerServer
	service service
}

type service interface {
	Get(context.Context, *api_model.GetUsersParams) ([]*model.User, error)
	GetByID(context.Context, int64) (*model.User, error)
	Create(context.Context, *api_model.CreateUserParams) (int64, error)
	Update(context.Context, int64, *api_model.UpdateUserParams) (*model.User, error)
	Delete(context.Context, int64) error
}

func (g *GrpcService) ListUsers(ctx context.Context, req *grpc.UserRequest) (*grpc.ListUserReply, error) {
	if req.Name == nil {
		return nil, status.Error(codes.InvalidArgument, "not input Name")
	}

	user := &api_model.GetUsersParams{
		Name: req.Name,
	}
	res, err := g.service.Get(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "test")
	}

	result := make([]*grpc.User, len(res))
	for _, item := range res {
		result = append(result, UserToGrpcUser(item))
	}

	return &grpc.ListUserReply{Users: result}, nil
}
