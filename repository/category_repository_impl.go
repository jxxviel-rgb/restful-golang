package repository

import (
	"context"
	"database/sql"
	"errors"
	"restful-golang/entity/model"
	"restful-golang/helper"
)

// f *file

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]model.Category, error) {
	SQL := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var categories []model.Category
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories, nil

}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (model.Category, error) {
	SQL := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category is not found")
	}
}
func (repository *CategoryRepositoryImpl) Store(ctx context.Context, tx *sql.Tx, category model.Category) (model.Category, error) {
	SQL := "insert into categories(name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category, nil
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category model.Category) (model.Category, error) {
	SQL := "update categories set name=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, category.Id, category.Name)
	helper.PanicIfError(err)
	return category, nil
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	SQL := "delete from categories where id=?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
}
