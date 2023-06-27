package repository

import (
	"context"
	"database/sql"
	"taxgo/16-uow/internal/db"
	"taxgo/16-uow/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dbt *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: dbt,
		Queries: db.New(dbt),
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}