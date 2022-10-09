package repository

import (
	"context"
	"database/sql"
	"restful-golang/entity/model"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]model.Category, error)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (model.Category, error)
	Store(ctx context.Context, tx *sql.Tx, category model.Category) (model.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category model.Category) (model.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
}
