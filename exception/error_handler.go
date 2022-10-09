package exception

import (
	"net/http"

	"github.com/jxxviel-rgb/restful-golang/entity/response"
	"github.com/jxxviel-rgb/restful-golang/helper"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {
	if errorNotFound(writer, request, error) {
		return
	}
	if validationErrors(writer, request, error) {
		return
	}
	internalServerError(writer, request, error)
}
func validationErrors(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToRequestBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
func errorNotFound(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := response.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToRequestBody(writer, webResponse)
		return true
	} else {
		return false
	}

}
func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := response.Response{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   error,
	}

	helper.WriteToRequestBody(writer, webResponse)
}
