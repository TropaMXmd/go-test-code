package main 

import(
	"fmt"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path ", r.URL.Path)
	fmt.Println("scheme ",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key: " ,k)
		fmt.Println("val: ",strings.Join(v, ""))
	}
	fmt.Fprintf(w,"Hello Tropa!")
}	

func main(){
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		panic(err)
	}
}

//http://localhost:9000/?url_long=111&name=tropa