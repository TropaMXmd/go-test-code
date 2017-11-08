package main 

import "fmt"

type person struct{
	fname string
	lname string
}
type secretAgent struct{
	person //get all the properties human struct have
	license bool
}
func (p person)speek(){
	fmt.Println("My name is ", p.fname)
}
func (s secretAgent)speek() {
	fmt.Println("I have license", s.license)
}
func main(){
	p := person{"Tropa","Mahmood"}
	p.speek()

	agent := secretAgent{p,true}
	agent.speek()
}