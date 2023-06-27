package repository

import (
	"context"
	"database/sql"
	"taxgo/16-uow/internal/db"
	"taxgo/16-uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dbt *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
