package main 

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
)

const(
	DB_HOST = "127.0.0.1"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASSWORD = "test123"
	DB_NAME = "moviesdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}

	err = db.DB().Ping()

	fmt.Println(err)

	db.Close()
}