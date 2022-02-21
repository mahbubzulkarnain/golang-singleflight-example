package psql

import "os"

type Config struct {
	Host string
	User string
	Pass string
	Port string
	Name string

	SslMode string
	Tz      string

	DSN string

	SSH *SSH
}

func NewConfig(ssh *SSH) Config {
	return Config{
		Host:    os.Getenv("DB_HOST"),
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		Port:    os.Getenv("DB_PORT"),
		Name:    os.Getenv("DB_NAME"),
		SslMode: os.Getenv("DB_SSLMODE"),
		Tz:      os.Getenv("DB_TZ"),
		SSH:     ssh,
	}
}
