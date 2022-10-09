package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jxxviel-rgb/restful-golang/helper"
	"github.com/jxxviel-rgb/restful-golang/service"

	req "github.com/jxxviel-rgb/restful-golang/entity/request"
	res "github.com/jxxviel-rgb/restful-golang/entity/response"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryControllerImpl(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := req.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := res.Response{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToRequestBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := req.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := res.Response{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToRequestBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	controller.CategoryService.Delete(request.Context(), id)
	webResponse := res.Response{
		Code:   200,
		Status: "OK",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	result := controller.CategoryService.FindById(request.Context(), id)
	// helper.PanicIfError(err)
	webResponse := res.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToRequestBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	result := controller.CategoryService.FindAll(request.Context())
	// helper.PanicIfError(err)
	webResponse := res.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToRequestBody(writer, webResponse)
}
