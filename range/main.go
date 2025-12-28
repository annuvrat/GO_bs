package main

import "fmt"

func main() {
 nums := []int{1,2,3,4,5}
	// for i:=0; i<len(nums); i++{
	// 	println(nums[i])
	// }
    sum:=0
	for _,num := range nums{
		sum =sum+num
	}
	fmt.Println(sum)
}
