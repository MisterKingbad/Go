package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Categories []Category `gorm:"many2many:products_categories"`
	Price      float64
	gorm.Model
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	category2 := Category{Name: "Banheiro"}
	db.Create(&category2)

	// db.Create(&Product{
	// 	Name:       "Note",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// })
	db.Create(&Product{
		Name:       "Note",
		Price:      1000.00,
		Categories: []Category{category},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range categories.Products {
			println("-", product.Name)
		}
	}

}
