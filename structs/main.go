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


func main() {

	myorder:=order{
		id: "1234",
		amount: 250.5,
		status: "created",
		// createdAt: time.Now(),
	}
  
	myorder.createdAt = time.Now()

fmt.Println("Order before update:", myorder)
}
