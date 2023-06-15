package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	CategoryID int
	Category   Category
	Price      float64
	SerialNumber SerialNumber
	gorm.Model
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
		Number: "123456",
		ProductID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range categories.Products {
			println("-", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}

}
