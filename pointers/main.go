package main

import "fmt"



// func number(num int){
// // in this case num is a copy of main function num
// // so changing num here will not affect main function num
// 	num =5
// 	fmt.Println("Inside number function:", num)

// }
// func main(){
// 	num:=1

// 	number(num)

// 	fmt.Println("Inside main function:", num)

// }
// lets now see how to use pointers to change the value of num in main function
// by reference

func changeNum(num*int){
	// here num is a pointer to main function num
	// so changing the value at the address will affect main function num
	*num =5 //dereferencing the pointer to change the value
	fmt.Println("Inside changeNum function:", *num)

}
func main(){
num:=1

changeNum(&num)

fmt.Println("Inside main function:", num)
// fmt.Println("memory address:", &num)

}