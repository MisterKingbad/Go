// go:build wireinject
// +build wireinject

package main

import (
	"taxgo/17-di/product"

	"github.com/google/wire"
	"database/sql"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}