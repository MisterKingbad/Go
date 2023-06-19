package main

import (
	"net/http"
	"taxgo/7-apis/configs"
	"taxgo/7-apis/internal/entity"
	"taxgo/7-apis/internal/infra/database"
	"taxgo/7-apis/internal/infra/webserver/handlers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProductById)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
		// r.Get("/products", productHandler.GetProductsPerPage)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.Create)
		r.Post("/login", userHandler.GetJWT) //generate_token
	})

	http.ListenAndServe(":8080", r)
}

