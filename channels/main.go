package main

import (
	"fmt"
	"time"
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
// func sum(result chan int,num1 int , num2 int){
// 	sum:= num1+num2

// 	result<-sum// blocking

// }

// func main(){

// 	result:= make(chan int)
// 	go sum(result,3,5)

// 	res:=<-result

// 	fmt.Println("Sum is:",res)
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

// }

//go routine synchronization using channels
// func task(done chan bool){
// defer func (){done<-true}()
// fmt.Println("task started")

// }

// func main(){

//  done:= make (chan bool)

//  go task(done)

//  <-done

// }

//buffered channels
// make(chan int)      // Unbuffered: Send BLOCKS until receive
// make(chan int, 100) // Buffered: Holds 100 values before blocking


func emailSender(emailChan chan string,done chan bool){

defer func(){done<-true}()

for email:=range emailChan{
	fmt.Println("sending email to ",email)
	time.Sleep(time.Second)
}


}

func main(){
	emailChan:= make(chan string,100)// buffered channel have the space to hold 100 messages
	
	done:=make(chan bool)

	go emailSender(emailChan,done)


	for i:=0;i<5;i++{

		emailChan<-fmt.Sprintf("%d@gmail.com",i)
	
	}
	

	fmt.Println("done sending")
	close(emailChan)
	<-done
}