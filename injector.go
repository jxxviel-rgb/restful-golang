//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"restful-golang/app"
	"restful-golang/controller"
	"restful-golang/middleware"
	"restful-golang/repository"
	"restful-golang/service"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepositoryImpl), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*repository.CategoryRepositoryImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedApp() *http.Server {
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
