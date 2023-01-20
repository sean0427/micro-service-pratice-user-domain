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
	if req.Name == nil || *req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "input name should not be nil")
	}

	user := &api_model.GetUsersParams{
		Name: req.Name,
	}
	res, err := g.service.Get(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
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

func (g *GrpcService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserReply, error) {
	if req.Id == nil || *req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "input id shuld not be nil")
	}

	res, err := g.service.GetByID(ctx, *req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if res == nil {
		return nil, status.Error(codes.NotFound, "")
	}

	return &pb.UserReply{User: UserToGrpcUser(res)}, nil
}

func (g *GrpcService) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.MsgReply, error) {
	if req.Id != nil {
		return nil, status.Error(codes.InvalidArgument, "input id shuld be nil")
	}
	if req.Name == nil || *req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "input name should not be nil")
	}

	user := &api_model.CreateUserParams{}
	if req.Name != nil {
		user.Name = *req.Name
	}

	res, err := g.service.Create(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.MsgReply{Id: api_model.Int64ToPointer(res)}, nil
}

func (g *GrpcService) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserReply, error) {
	if req.Id == nil || *req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "input id shuld not be nil")
	}

	user := &api_model.UpdateUserParams{}
	// TODO: other pramas
	if req.Name != nil {
		user.Name = *req.Name
	}

	res, err := g.service.Update(ctx, *req.Id, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if res == nil {
		return nil, status.Error(codes.NotFound, "")
	}

	return &pb.UserReply{User: UserToGrpcUser(res)}, nil
}

func (g *GrpcService) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.MsgReply, error) {
	if req.Id == nil || *req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "input id shuld not be nil")
	}

	err := g.service.Delete(ctx, *req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.MsgReply{Id: req.Id}, nil
}
