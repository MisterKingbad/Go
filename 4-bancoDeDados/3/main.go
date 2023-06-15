package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	Price        float64
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Note",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// })
	db.Create(&Product{
		Name:       "Note",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	db.Create(&SerialNumber{
		Number:    "000000",
		ProductID: 2,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}
