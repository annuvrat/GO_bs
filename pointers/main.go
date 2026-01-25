package main

import "fmt"



func number(num int){

	num =5
	fmt.Println("Inside number function:", num)

}
func main(){
	num:=1

	number(num)

	fmt.Println("Inside main function:", num)

}
