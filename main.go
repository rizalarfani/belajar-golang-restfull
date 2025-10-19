package main

import (
	"log"
	"rizalarfani/belajar-restful-api/helper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := InitializedServer()

	log.Printf("starting HTTP server on %s", ":3000")
	if err := server.Run(":3000"); err != nil {
		helper.PanicIfError(err)
	}
}
