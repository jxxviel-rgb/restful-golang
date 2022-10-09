package service

import (
	"context"
	"restful-golang/entity/request"
	"restful-golang/entity/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
