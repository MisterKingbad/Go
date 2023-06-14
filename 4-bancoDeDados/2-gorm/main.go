package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	// db.Create(&Product{
	// 	Name: "Pc da nasa",
	// 	Price: 100000.00,
	// })
	// products := []Product{
	// 	{
	// 		Name: "Pc da nasa",
	// 		Price: 100000.00,
	// 	},
	// 	{
	// 		Name: "Note do Elon musk",
	// 		Price: 1000000000000000000000000.00,
	// 	},
	// 	{
	// 		Name: "tesla",
	// 		Price: 1000000.00,
	// 	},
	// }
	// db.Create(&products)

	// Select one
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "tesla")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }
		


}
