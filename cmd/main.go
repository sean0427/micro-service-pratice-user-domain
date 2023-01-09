package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	service "github.com/sean0427/micro-service-pratice-user-domain"
	config "github.com/sean0427/micro-service-pratice-user-domain/config"
	handler "github.com/sean0427/micro-service-pratice-user-domain/net"
	repository "github.com/sean0427/micro-service-pratice-user-domain/postgres"
)

func NewSQLDB() (*gorm.DB, error) {
	dsn, err := config.GetPostgresDNS()
	if err != nil {
		return nil, err
	}

	gormConfig := &gorm.Config{}
	conn, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func startServer() {
	fmt.Println("Starting server...")

	db, err := NewSQLDB()
	if err != nil {
		panic(err)
	}

	r := repository.New(db)
	s := service.New(r)
	h := handler.New(s)

	handler := h.InitHandler()
	http.ListenAndServe(":8080", handler)

	fmt.Println("Stoping server...")
}

func main() {
	startServer()
}
