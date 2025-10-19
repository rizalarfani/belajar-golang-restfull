package controller

import (
	"encoding/json"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/model/web"
	"rizalarfani/belajar-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(c *gin.Context) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(c.Request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(c.Request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	categoryId := c.Params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(c.Request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(c *gin.Context) {
	categoryId := c.Params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(c.Request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(c *gin.Context) {
	categoryId := c.Params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(c.Request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(c *gin.Context) {
	categoryResponses := controller.CategoryService.FindAll(c.Request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}
