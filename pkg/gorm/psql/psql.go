package psql

import (
	"database/sql"
	"errors"
	"fmt"

	"net"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
)

type PSQL struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock

	net    net.Conn
	ssh    *ssh.Client
	config Config
}

func Open(config Config) (conn *PSQL, err error) {
	conn = &PSQL{}

	defer func() {
		if err != nil && conn != nil {
			conn.Close()
		}
	}()

	if config.SSH != nil {
		conn.net, conn.ssh, err = config.SSH.NewConnection()
		if err != nil {
			return nil, err
		}
	}

	conn.db, err = config.NewConnection()
	return
}

func Mock() (conn *PSQL, err error) {
	conn = &PSQL{}
	conn.db, conn.mock, err = NewConfig(nil).NewConnectionMock()
	return
}

func (f PSQL) SQL() *sql.DB {
	if f.Gorm() == nil {
		panic(errors.New("missing connection"))
	}
	sqlDB, err := f.Gorm().DB()
	if err != nil {
		panic(err)
	}
	return sqlDB
}

func (f PSQL) Gorm() *gorm.DB {
	return f.db
}

func (f PSQL) Mock() sqlmock.Sqlmock {
	return f.mock
}

func (f PSQL) Close() {
	var err error
	if f.ssh != nil {
		if err = f.ssh.Close(); err != nil {
			fmt.Println(err)
		}
	}

	if f.net != nil {
		if err = f.net.Close(); err != nil {
			fmt.Println(err)
		}
	}

	var sqlDB *sql.DB
	if sqlDB, err = f.Gorm().DB(); err == nil {
		err = sqlDB.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
}
