package service

import (
	"context"

	userModel "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user"
)

func (s service) Find(ctx context.Context) (userModel.Users, error) {
	return s.userRepository.Find(ctx)
}
