package main

import (
	"context"
	"database/sql"
	"taxgo/15-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql","root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID: uuid.New().String(),
	// 	Name: "BAckend1",
	// 	Description: sql.NullString{String: "Back description1", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID: "95820388-9723-4f1c-8e0a-5d31d12cefc1",
	// 	Name: "Back [UPD]",
	// 	Description: sql.NullString{String: "backUp", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	err = queries.DeleteCategory(ctx, "95820388-9723-4f1c-8e0a-5d31d12cefc1")

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}