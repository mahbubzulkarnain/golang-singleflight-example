package repository

import (
	"gorm.io/gorm"

	userRepository "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user/repository"
)

type V1 struct {
	UserRepository userRepository.Repository
}

type Repository struct {
	V1
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		V1{
			UserRepository: userRepository.NewRepository(db),
		},
	}
}
