package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jxxviel-rgb/restful-golang/app"
	"github.com/jxxviel-rgb/restful-golang/controller"
	"github.com/jxxviel-rgb/restful-golang/entity/model"
	"github.com/jxxviel-rgb/restful-golang/helper"
	"github.com/jxxviel-rgb/restful-golang/middleware"
	"github.com/jxxviel-rgb/restful-golang/repository"
	"github.com/jxxviel-rgb/restful-golang/service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// setting up router for testing
func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryControllerImpl(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

// setting up DB for testing
func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:fikri@tcp(localhost:3306)/belajarrestful_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

// here a test begin
func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	categoryRepository := repository.NewCategoryRepositoryImpl()

	tx, _ := db.Begin()
	category, _ := categoryRepository.Store(context.Background(), tx, model.Category{
		Name: "Gadget",
	})
	tx.Commit()
	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
func TestUpdateFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	categoryRepository := repository.NewCategoryRepositoryImpl()

	tx, _ := db.Begin()
	category, _ := categoryRepository.Store(context.Background(), tx, model.Category{
		Name: "Gadget",
	})
	tx.Commit()
	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	categoryRepository := repository.NewCategoryRepositoryImpl()

	tx, _ := db.Begin()
	category, _ := categoryRepository.Store(context.Background(), tx, model.Category{
		Name: "Gadget",
	})
	tx.Commit()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}
func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}
func TestListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	categoryRepository := repository.NewCategoryRepositoryImpl()
	tx, _ := db.Begin()
	category1, _ := categoryRepository.Store(context.Background(), tx, model.Category{
		Name: "Gadget",
	})
	category2, _ := categoryRepository.Store(context.Background(), tx, model.Category{
		Name: "Computer",
	})
	tx.Commit()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	categoryRepository := repository.NewCategoryRepositoryImpl()

	tx, _ := db.Begin()
	category, _ := categoryRepository.Store(context.Background(), tx, model.Category{
		Name: "Gadget",
	})
	tx.Commit()
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}
func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestUnautorized(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ASD")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

}
