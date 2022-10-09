package helper

import (
	"github.com/jxxviel-rgb/restful-golang/entity/model"
	"github.com/jxxviel-rgb/restful-golang/entity/response"
)

func ToCategoryResponse(category model.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func ToCategoryResponses(categories []model.Category) []response.CategoryResponse {
	var categoriesResponse []response.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}
