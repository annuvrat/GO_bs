package main

import "fmt"


func printslice[T comparable, V string](items []T, name V) {

	for _, item := range items {
		println( item,name)
	}

}

type stack[T comparable]struct{
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
   printslice([]int {1,2,3}, "numbers")
	
}
