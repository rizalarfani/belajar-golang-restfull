package main

import (
	"net/http"
	"rizalarfani/belajar-restful-api/app"
	"rizalarfani/belajar-restful-api/controller"
	"rizalarfani/belajar-restful-api/exception"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/middleware"
	"rizalarfani/belajar-restful-api/repository"
	"rizalarfani/belajar-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryControler := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryControler.FindAll)
	router.GET("/api/categories/:categoryId", categoryControler.FindById)
	router.POST("/api/categories", categoryControler.Create)
	router.PUT("/api/categories/:categoryId", categoryControler.Update)
	router.DELETE("/api/categories/:categoryId", categoryControler.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
