package service

import (
	"github.com/mahbubzulkarnain/golang-singleflight-example/internal/repository"
	userService "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user/service"
)

type V1 struct {
	UserService userService.Service
}

type Service struct {
	V1
}

func NewService(r repository.Repository) Service {
	return Service{
		V1{
			UserService: userService.NewService(r.V1.UserRepository),
		},
	}
}
