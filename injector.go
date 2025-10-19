//go:build wireinject
// +build wireinject

package main

import (
	"rizalarfani/belajar-restful-api/app"
	"rizalarfani/belajar-restful-api/controller"
	"rizalarfani/belajar-restful-api/repository"
	"rizalarfani/belajar-restful-api/routes"
	"rizalarfani/belajar-restful-api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var routerSet = wire.NewSet(
	routes.NewRoutesConfig,
	wire.Struct(new(routes.Routes), "*"),
)

func InitializedServer() *gin.Engine {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		routerSet,
		app.NewRouter,
	)
	return nil
}
