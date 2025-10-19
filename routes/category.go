package routes

import (
	"rizalarfani/belajar-restful-api/controller"
	"rizalarfani/belajar-restful-api/middleware"

	"github.com/gin-gonic/gin"
)

type CategoryRoute struct {
	controller controller.CategoryController
	group      *gin.RouterGroup
}

type ICategoryRoute interface {
	Run()
}

func NewCategoryRoute(controller controller.CategoryController, group *gin.RouterGroup) *CategoryRoute {
	return &CategoryRoute{
		controller: controller,
		group:      group,
	}
}

func (category *CategoryRoute) Run() {
	category.group.GET("/categories", middleware.ValidateApiKey(), category.controller.FindAll)
	category.group.GET("/categories/:categoryId", middleware.ValidateApiKey(), category.controller.FindById)
	category.group.POST("/categories", middleware.ValidateApiKey(), category.controller.Create)
	category.group.PUT("/categories/:categoryId", middleware.ValidateApiKey(), category.controller.Update)
	category.group.DELETE("/categories/:categoryId", middleware.ValidateApiKey(), category.controller.Delete)
}
