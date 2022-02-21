package service

import (
	"context"

	userModel "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user"
	userRepository "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user/repository"
)

type Service interface {
	Find(ctx context.Context) (userModel.Users, error)
}

type service struct {
	userRepository userRepository.Repository
}

func NewService(userRepository userRepository.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}
