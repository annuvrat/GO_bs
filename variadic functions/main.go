package main

import "fmt"



func sum(nums...int)int{
	total:= 0
	for _,num:= range nums{
		total+=num
	}
	return total
}


func main(){
  
   nums:=[]int{2,3,4,5,6}
   result:= sum(nums...)
	// result:= sum(2,3,4,5,6)

	fmt.Println("The sum is:",result)
 
}