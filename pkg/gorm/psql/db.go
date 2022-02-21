package psql

import (
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (c Config) NewConnectionMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))

	return db, mock, nil
}

func (c Config) NewConnection() (*gorm.DB, error) {
	if c.DSN == "" {
		c.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", c.Host, c.User, c.Pass, c.Name, c.Port, c.SslMode, c.Tz)
	}

	dialector := postgres.Open(c.DSN)
	if c.SSH != nil {
		dialector = postgres.New(postgres.Config{
			DriverName: "postgres+ssh",
			DSN:        c.DSN,
		})
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	d, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	d.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	d.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	d.SetConnMaxLifetime(time.Hour)

	if err = d.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
