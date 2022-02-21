package main

import (
	"github.com/mahbubzulkarnain/golang-singleflight-example/internal/delivery/http"
	"github.com/mahbubzulkarnain/golang-singleflight-example/internal/repository"
	"github.com/mahbubzulkarnain/golang-singleflight-example/internal/service"
	"github.com/mahbubzulkarnain/golang-singleflight-example/pkg/gorm/psql"
)

func main() {
	db, err := psql.Mock()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := repository.NewRepository(db.Gorm())
	s := service.NewService(r)
	http.RegisterHTTPServer("3000", s)
}
