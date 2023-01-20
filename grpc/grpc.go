package grpc

import (
	"context"

	"github.com/sean0427/micro-service-pratice-user-domain/api_model"
	pb "github.com/sean0427/micro-service-pratice-user-domain/grpc/grpc"
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
	pb.UnimplementedUserHandlerServer
	service service
}

type service interface {
	Get(context.Context, *api_model.GetUsersParams) ([]*model.User, error)
	GetByID(context.Context, int64) (*model.User, error)
	Create(context.Context, *api_model.CreateUserParams) (int64, error)
	Update(context.Context, int64, *api_model.UpdateUserParams) (*model.User, error)
	Delete(context.Context, int64) error
}

func (g *GrpcService) ListUsers(ctx context.Context, req *pb.UserRequest) (*pb.ListUserReply, error) {
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

	if len(res) == 0 {
		status.Error(codes.NotFound, "not found")
	}

	result := make([]*pb.User, len(res))
	for i, item := range res {
		result[i] = UserToGrpcUser(item)
	}

	return &pb.ListUserReply{Users: result}, nil
}
