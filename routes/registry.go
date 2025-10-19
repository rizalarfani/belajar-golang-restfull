package routes

import (
	"log"
	"rizalarfani/belajar-restful-api/controller"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Prefix string
}

func NewRoutesConfig() Config {
	return Config{
		Prefix: "/api/v1",
	}
}

type Routes struct {
	Config   Config
	Category controller.CategoryController
}

func (r *Routes) Mount(e *gin.Engine) {
	if e == nil {
		log.Fatal("gin.Engine is nil in Routes.Mount")
	}
	log.Printf("[mount] prefix=%q\n", r.Config.Prefix)

	api := e.Group(r.Config.Prefix)
	NewCategoryRoute(r.Category, api).Run()

	for _, rt := range e.Routes() {
		log.Printf("[route] %s %s", rt.Method, rt.Path)
	}
}
