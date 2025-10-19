package app

import (
	"log"
	"net/http"
	"rizalarfani/belajar-restful-api/exception"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/model/web"
	"rizalarfani/belajar-restful-api/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter(reg *routes.Routes) *gin.Engine {
	router := gin.Default()
	router.Use(exception.ErrorHandler())

	router.NoRoute(func(ctx *gin.Context) {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not FOUND",
		}
		helper.WriteToResponseBody(ctx.Writer, webResponse)
	})

	router.GET("/", func(ctx *gin.Context) {
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Welcome to api with golang",
		}
		helper.WriteToResponseBody(ctx.Writer, webResponse)
	})

	if reg == nil {
		log.Fatal("routes.Routes (reg) is nil in NewRouter (Wire tidak memprovide reg)")
	}

	reg.Mount(router)
	return router
}
