package main 

import(
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct{
	Stuff Data
}

type Data struct{
	Fruit Fruits
	Vegetable Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request){
	response,err := getJsonResponse()
	if err != nil{
		panic(err)
	}
	fmt.Fprintf(w,string(response))
}

func main() {
	http.HandleFunc("/",serveRest)
	http.ListenAndServe("localhost:1337",nil)
}
func getJsonResponse()([]byte,error){
	fruits := make(map[string]int)
	fruits["orange"] = 22
	fruits["mango"] = 5

	veges := make(map[string]int)
	veges["tomato"] = 10
	veges["chili"] = 20

	d := Data{fruits,veges}
	p := Payload{d}

	return json.MarshalIndent(p,""," ")
}
