package main

import "fmt"


func printslice[T int|string|bool](items []T) {

	for _, item := range items {
		println(item)
	}

}

type stack[T any]struct{
	elements[]T
}
func main() {


	mystack:= stack[string]{
		elements: []string{"first","second","third"},
	}
	// nums:=[]int{1,2,3,4,5}

	// names:= []string{"Alice", "Bob", "Charlie"}

	// vals:=  []bool{true,false}

   fmt.Println(mystack)  
	
}
