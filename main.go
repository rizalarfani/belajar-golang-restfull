package main

import (
	"log"
	"net/http"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()

	log.Printf("starting HTTP server on %s", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
