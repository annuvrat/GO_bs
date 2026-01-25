package main

import (
	"fmt"
	"time"
)


type order struct{
	id string
	amount float32
	status string
	createdAt time.Time  //nano second precision
}

func (o *order)changeStatus(status string){
	o.status=status

}

func (o order)getAmount()float32{

	return o.amount

}
func main() {

	myorder:=order{
		//if you  dont initialize fields, they get zero or default values and for string its ""
		//int =0
		//string=""
		//float=0.0
		//bool=false
		id: "1234",
		// amount: 250.5,
		status: "created",
		// createdAt: time.Now(),
	}


	MYorder2:= order{
		id: "5678",
		amount: 100.0,
		// status: "shipped",
		createdAt: time.Now(),
	}
    myorder.changeStatus("confirmed")
	// fmt.Println("Order after changeStatus call:", myorder)
	myorder.createdAt = time.Now()
	// MYorder2.status = "delivered"
	fmt.Println("Order Amount:", myorder.getAmount())

// fmt.Println("Order before update:", myorder)
fmt.Println("Second Order:", MYorder2)
// fmt.Println("Order created at:", myorder.createdAt)
}
