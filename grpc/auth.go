package grpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/sean0427/micro-service-pratice-user-domain/grpc/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authSerivice interface {
	Authenticate(ctx context.Context, name, password string) (bool, error)
}

type AuthGrpc struct {
	pb.UnimplementedAuthServer
	service authSerivice
}

func NewAuthHandler(service authSerivice) *AuthGrpc {
	return &AuthGrpc{
		service: service,
	}
}

func (g *AuthGrpc) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	if req == nil || req.Name == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "auth should have name and password")
	}

	success, err := g.service.Authenticate(ctx, req.Name, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.AuthReply{
		Message: fmt.Sprintf("authenticated timestamp: %v", time.Now()),
		Success: success,
	}, nil
}
