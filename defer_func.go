package main 

import "fmt"

func hello(){
	fmt.Println("Hello")
}
func world(){
	fmt.Println("World")
}
func main() {
	defer world() //run this at the last possible moment before the end of execution of main function
	hello()
}