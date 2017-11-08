package main 

import "fmt"

type human struct{
	fname string
	lname string
}
func (h human)speak(){
	fmt.Println("My name is ",h.fname)
}
func main(){
	h := human{"Tropa","Mahmood"}
	h.speak()
}