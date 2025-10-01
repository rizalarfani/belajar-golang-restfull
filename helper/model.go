package helper

import (
	"rizalarfani/belajar-restful-api/model/domain"
	"rizalarfani/belajar-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   int(category.Id),
		Name: category.Name,
	}
}
