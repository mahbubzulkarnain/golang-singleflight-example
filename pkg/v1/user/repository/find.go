package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	userModel "github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user"
)

var userRepositoryFindGroup singleflight.Group

func (r repository) Find(ctx context.Context) (users userModel.Users, err error) {
	sql := r.DB.WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&users)
	})

	data, err, shared := userRepositoryFindGroup.Do(sql, func() (interface{}, error) {
		time.Sleep(time.Second * 5)

		var data userModel.Users
		if err = faker.FakeData(&data); err != nil {
			return nil, err
		}
		return data, nil

		//if err = r.DB.Raw(sql).Find(&data).Error; err != nil {
		//	return nil, err
		//}
		//return data, nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("shared?", shared)

	var ok bool
	if users, ok = data.(userModel.Users); !ok {
		return nil, errors.New("something wrong")
	}
	return
}
