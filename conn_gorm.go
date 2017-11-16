package main 

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_"github.com/lib/pq"
)

const(
	DB_HOST = "127.0.0.1"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASSWORD = "test123"
	DB_NAME = "moviesdb"
)

type Product struct{
	gorm.Model
	Code string
	Price uint
}

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
  	db.AutoMigrate(&Product{})

  	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	fmt.Println(&product)
}