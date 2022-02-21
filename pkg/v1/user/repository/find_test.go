package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mahbubzulkarnain/golang-singleflight-example/pkg/v1/user"
)

func Test_repository_Find(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()
	defer mock.ExpectClose()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	require.NoError(t, err)

	var rowsData = make(user.Users, 0)
	require.NoError(t, faker.FakeData(&rowsData))

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		registerMock func(mock sqlmock.Sqlmock, result user.Users)
	}{
		{
			name: "should success",
			fields: fields{
				DB: gormDB,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			registerMock: func(mock sqlmock.Sqlmock, users user.Users) {
				expectedSQL := regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)

				columns := []string{"id", "firstname", "lastname", "email", "created_at", "updated_at", "deleted_at"}
				rows := sqlmock.NewRows(columns)
				for _, u := range users {
					rows.AddRow(u.ID, u.Firstname, u.Lastname, u.Email, u.CreatedAt, u.UpdatedAt, u.DeletedAt)
				}
				mock.ExpectQuery(expectedSQL).WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.registerMock(mock, rowsData)

			r := repository{
				DB: tt.fields.DB,
			}
			_, err := r.Find(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
