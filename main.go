package main

import (
	"net/http"
	"restful-golang/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	// server := InitializedApp()

	// err := server.ListenAndServe()
	// helper.PanicIfError(err)

}
