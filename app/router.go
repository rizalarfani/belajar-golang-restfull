package app

import (
	"rizalarfani/belajar-restful-api/controller"
	"rizalarfani/belajar-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryControler controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryControler.FindAll)
	router.GET("/api/categories/:categoryId", categoryControler.FindById)
	router.POST("/api/categories", categoryControler.Create)
	router.PUT("/api/categories/:categoryId", categoryControler.Update)
	router.DELETE("/api/categories/:categoryId", categoryControler.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
