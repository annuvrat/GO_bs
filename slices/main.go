package main

import (
	"fmt"
	"slices"
)

func main(){
	// var nums [3]int
//   fmt.Println(nums)


// var nums = make([]int,2,5)
// //cap= capacity

// nums = append(nums,1)
// nums = append(nums,2)
// nums = append(nums,3)
// nums = append(nums,4)
// nums = append(nums,5)
// var nums2 = make([]int,len(nums))
// // fmt.Println(cap(nums))
// nums = append(nums,2)
// fmt.Println(nums,nums2)
// fmt.Println(cap(nums))
var nums = []int{1,2,3,4,5}
var num2 = []int{1,2,3,4,5}
fmt.Println(slices.Equal(nums,num2))

}