package repository

import (
	"context"
	"gorm.io/gorm"

	userModel "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user"
)

type Repository interface {
	Find(ctx context.Context) (users userModel.Users, err error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
