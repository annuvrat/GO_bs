package main

import (
	"fmt"
	// "math/rand"
	// "time"
)
// sending example
// func processNum(numChan chan int){
// for num:=range numChan{
// 	fmt.Println("processing number",num)
// 	time.Sleep(time.Second*1)
// }
// }
func sum(result chan int,num1 int , num2 int){
	sum:= num1+num2

	result<-sum

}


func main(){


	result:= make(chan int)
	go sum(result,3,5)

	res:=<-result
	
	fmt.Println("Sum is:",res)
// 	numChan := make(chan int)

// 	go processNum(numChan)

//    for {
// 	numChan <- rand.Intn(100)
//    }
	// time.Sleep(time.Second*2)
// 	messageChan:= make(chan string)

//   messageChan<-"ping"//blocking

//   msg:=<-messageChan

//   fmt.Println(msg)

}
