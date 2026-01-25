package main

import (
	"fmt"
	"time"
	// "time"
)

// type order struct{
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time  //nano second precision
// }

// func newOrder(id string, amount float32, status string)*order{
// myOrder:= order{
// 	id: id,
// 	amount: amount,
// 	status: status,
// }
// return &myOrder
// }
// func (o *order)changeStatus(status string){
// 	o.status=status

// }

// func (o order)getAmount()float32{

// 	return o.amount

// }
// func main() {

// 	myorder:=order{
// 		//if you  dont initialize fields, they get zero or default values and for string its ""
// 		//int =0
// 		//string=""
// 		//float=0.0
// 		//bool=false
// 		id: "1234",
// 		// amount: 250.5,
// 		status: "created",
// 		// createdAt: time.Now(),
// 	}

// 	MYorder2:= order{
// 		id: "5678",
// 		amount: 100.0,
// 		// status: "shipped",
// 		createdAt: time.Now(),
// 	}
//     myorder.changeStatus("confirmed")
// 	// fmt.Println("Order after changeStatus call:", myorder)
// 	myorder.createdAt = time.Now()
// 	// MYorder2.status = "delivered"
// 	fmt.Println("Order Amount:", myorder.getAmount())

// // fmt.Println("Order before update:", myorder)
// fmt.Println("Second Order:", MYorder2)
// // fmt.Println("Order created at:", myorder.createdAt)
// }

// func  main() {

//   language:=struct{
// 	name string
// 	year int
//   }{
// 	name: "Go",
// 	year: 2009,
//   }
//   fmt.Println("Language:", language)
// }


type customer struct{
	id string
	name string
	email string
}


type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
	customer

}


func main(){

	result:= order{
		id: "ord123",
		amount: 299.99,
		status: "processing",
	}

	// result.customer.email="email"
	// result.customer.id="cust001"


	result.customer= customer{
		id: "cust001",
		name: "John Doe",
		email: "empr",
	}
fmt.Println("Order:", result)
}