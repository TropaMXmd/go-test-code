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

type Owner struct{
	gorm.Model
	FirstName string
	LastName string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	//db.SingularTable(true) //Defeat the pluralization algorithm in table name

	db.DropTableIfExists(&Owner{})
	db.CreateTable(&Owner{})


	owner := Owner{
		FirstName: "Joe",
		LastName: "Smith",
	}

	db.Create(&owner) //insert data in db by sending a reference to the owner object
	owner.FirstName = "Tropa"

	//db.Save(&owner) //Update the owner data if exist otherwise create owner data

	db.Save(&owner) //Print the sql log command
	
	var o Owner
	db.Debug().First(&o)
	fmt.Println(o)

	o = Owner{}
	db.Delete(&owner) //Delete data -> set deleted_at column & print log
	db.Debug().First(&o)
	fmt.Println(o)

}

//add a method to the struct that returns the tablename
// func (o *Owner) TableName() string{
// 	return "foo"
// }