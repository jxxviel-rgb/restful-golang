//go:build wireinject
// +build wireinject

package main

import (
	"github.com/jxxviel-rgb/restful-golang/app"

	"github.com/jxxviel-rgb/restful-golang/controller"
	"github.com/jxxviel-rgb/restful-golang/middleware"
	"github.com/jxxviel-rgb/restful-golang/repository"
	"github.com/jxxviel-rgb/restful-golang/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// var categorySet = wire.NewSet(
// 	repository.NewCategoryRepositoryImpl,
// 	wire.Bind(new(repository.CategoryRepositoryImpl), new(*repository.CategoryRepositoryImpl)),
// 	service.NewCategoryServiceImpl,
// 	wire.Bind(new(service.CategoryService), new(*repository.CategoryRepositoryImpl)),
// 	controller.NewCategoryControllerImpl,
// 	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
// )

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
