package main


import "fmt"


// type OrderStatus string
type OrderStatus int

const (
	Pending OrderStatus = iota
	Shipped
	Delivered
	Canceled
)

// const (
// 	Pending OrderStatus = "Pending"
// 	Shipped OrderStatus = "Shipped"
// 	Delivered OrderStatus = "Delivered"
// 	Canceled OrderStatus = "Canceled"
// )


func Print(status ...OrderStatus){

	fmt.Println("Status is:",status)
}


func main(){

	Print(Canceled,Delivered,Pending)
}