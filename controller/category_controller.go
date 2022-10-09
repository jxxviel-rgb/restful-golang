package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Store(writer http.ResponseWriter, requst *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, requst *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, requst *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, requst *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, requst *http.Request, params httprouter.Params)
	
}
