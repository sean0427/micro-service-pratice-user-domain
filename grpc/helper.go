package grpc

import (
	"github.com/sean0427/micro-service-pratice-user-domain/grpc/grpc"
	"github.com/sean0427/micro-service-pratice-user-domain/model"
)

func UserToGrpcUser(item *model.User) *grpc.User {
	if item == nil {
		return nil
	}
	return &grpc.User{
		Id:    item.ID,
		Name:  item.Name,
		Email: item.Email,
	}
}
