package main 

import(
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"net/http"
	"encoding/json"
)

var db *sql.DB

const(
	DB_HOST = "127.0.0.1"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASSWORD = "test123"
	DB_NAME = "moviesdb"
)

// func init() {
// 	var err error
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 	"password=%s dbname=%s sslmode=disable",
// 	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

// 	db, err = sql.Open("postgres", psqlInfo)
// 	checkErr(err)
// 	defer db.Close()
// }

func getMovies(w http.ResponseWriter, r *http.Request){
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = sql.Open("postgres", psqlInfo)
	checkErr(err)
	defer db.Close()

	rows,err := db.Query("SELECT * FROM movieschema.movies")
	checkErr(err)
	defer rows.Close()
	columns, err := rows.Columns()
	checkErr(err)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
	  for i := 0; i < count; i++ {
	      valuePtrs[i] = &values[i]
	  }
	  rows.Scan(valuePtrs...)
	  entry := make(map[string]interface{})
	  for i, col := range columns {
	      var v interface{}
	      val := values[i]
	      b, ok := val.([]byte)
	      if ok {
	          v = string(b)
	      } else {
	          v = val
	      }
	      entry[col] = v
	  }
	  tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	checkErr(err)

	fmt.Fprintf(w,string(jsonData))
}
func storeMovie(w http.ResponseWriter, r *http.Request){
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = sql.Open("postgres", psqlInfo)
	checkErr(err)
	defer db.Close()

	r.ParseForm()                     // Parses the request body
    name := r.Form.Get("name") // x will be "" if parameter is not set
    genre := r.Form.Get("genre")

    var	lastInsertId int
	err	=	db.QueryRow("INSERT	INTO movieschema.movies(name,genre)	VALUES($1,$2)	returning	id;", name,genre).Scan(&lastInsertId)
	checkErr(err)

	fmt.Fprintf(w,"Data is inserted")
}
func removeMovie(w http.ResponseWriter, r *http.Request){
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = sql.Open("postgres", psqlInfo)
	checkErr(err)
	defer db.Close()

	r.ParseForm()                     // Parses the request body
    id := r.Form.Get("id") // x will be "" if parameter is not set
	stmt,err :=	db.Prepare("delete	from movieschema.movies	where	id=$1")
	checkErr(err)

	res,err	:= stmt.Exec(id)
	checkErr(err)

	affect,	err	:=	res.RowsAffected()
	checkErr(err)
	fmt.Println(affect,	"rows	changed")

	fmt.Fprintf(w,"Data is deleted successfully.")
}
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello! this is index page.")
}
func main(){
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/movies/store",storeMovie)
	http.HandleFunc("/movies/remove",removeMovie)
	http.HandleFunc("/",index)
	http.ListenAndServe(":9000",nil)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}