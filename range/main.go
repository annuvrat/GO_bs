package main

import "fmt"

func main() {
//  nums := []int{1,2,3,4,5}
	// for i:=0; i<len(nums); i++{
	// 	println(nums[i])
	// }
    // sum:=0
	
	// for i,num := range nums{
	// 	// fmt.Println(i)
	// 	fmt.Println(num,i)
	// }
	// k:= map[string]int{"price":100,"age":22}

	// for key,value:= range k{
	// 	fmt.Println(key,value)
	// }

// value will be a unicode code point
	for i,c:= range "hello"{
		fmt.Println(i,c)
	}
}
