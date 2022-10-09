package app

import (
	"github.com/jxxviel-rgb/restful-golang/controller"
	"github.com/jxxviel-rgb/restful-golang/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.POST("/api/categories", categoryController.Store)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.PanicHandler = exception.ErrorHandler

	return router
}
