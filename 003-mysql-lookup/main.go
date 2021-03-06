package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("mysql", "root:welcome@tcp(localhost:3306)/Demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("Database failed")
	}

	db.AutoMigrate(&Product{})

	var product Product
	db.First(&product, 1)

	fmt.Println(product)

	fmt.Println()
	os.Exit(0)
}
