package main 

import "fmt"

type testInt func(int) bool

func isOdd(x int) bool{
	if x % 2 == 0 {
		return false
	}
	return true
}

func isEven(x int) bool{
	if x%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _ ,value := range slice{
		if f(value) {
			result = append(result,value)
		}
	}

	return result
}

func main() {
	slice := []int{1,2,3,4,5,7}
	fmt.Println("slice = ",slice)
	odd := filter(slice,isOdd)
	fmt.Println("Odd slice = ",odd)
	even := filter(slice,isEven)
	fmt.Println("Even slice = ",even)		
}