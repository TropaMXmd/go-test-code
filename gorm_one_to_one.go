package main 

//One to one relationship: user(has one), creaditcard(belongs to) : 
//a user has a credit card, a credit card is belongs to a user

import(
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

type User struct{
	gorm.Model
	Name string
	CreditCard CreditCard
}

type CreditCard struct{
	gorm.Model
	UserId uint
	Number string
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

	//db.SingularTable(true) //Defeat the pluralization algorithm in table name

	db.AutoMigrate(&User{},&CreditCard{})

	user := User{
		Name: "Tropa",
		CreditCard: CreditCard{
			Number: "1234567",
		},
	}

	db.Create(&user)

	var card CreditCard

	db.Model(&user).Related(&card)
	fmt.Println(card)

	card = CreditCard{}
	//db.Delete(user)
	db.Model(&user).Debug().Delete(&card)

	db.Model(&user).Related(&card)
	fmt.Println(card)
}

