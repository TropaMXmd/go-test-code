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
func (p person)speak(){
	fmt.Println("My name is ", p.fname)
}
func (s secretAgent)speak() {
	fmt.Println("I have license", s.license)
}
type human interface{
	speak()
}
func saysomething(h human) {
	h.speak()
}
func main(){
	p := person{"Tropa","Mahmood"}
	p.speak()
	saysomething(p)

	agent := secretAgent{p,true}
	agent.speak()
	saysomething(agent)
}

//ekhne both secretagent person er sob properties pabe. 
//ekhne person and secretagent duitar e duita speek function ache.
//abar ekhne bola ache human interface er ekta function hobe speak jar return type ni and arg o ni
//so jar jar speak function ache tarao human type er........means o k implement korche
//human interface er ekta function hocche saysomething....jehetu person and secretagent human k implement korse so orao saysomthing function k access krte parbe
