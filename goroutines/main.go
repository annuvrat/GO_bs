package main

// goroutines are light wieght threads managed by go runtime
import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("Task",id,"is being processed")
}

func main(){
  for i :=0;i<=10;i++{
	// go task(i)

 go	func (i int){
		task(i)
	}(i)
  }
 

  time.Sleep(time.Second*2)
	// task(1)

}

