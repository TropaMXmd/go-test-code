package main 

import "fmt"

type Human struct{
	fname string
	age int
}
type Student struct{
	Human //Inherites human. so can access all functions & properties of Human
	std_id int
}
type Employee struct{
	Human
	emp_id int
}

func (h *Human)sayHi(){
	fmt.Println("Say hi to ",h.fname)
}
func main() {
	st := Student{Human{"Tropa",26}, 23}
	st.sayHi()
}