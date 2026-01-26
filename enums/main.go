package main


import "fmt"


type OrderStatus int

const (
	Pending OrderStatus = iota
	Shipped
	Delivered
	Canceled
)


func Print(status OrderStatus){

	fmt.Println("Status is:",status)
}


func main(){

	Print(Canceled)
}