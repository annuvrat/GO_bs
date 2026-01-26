package main

import "sync"

// goroutines are light wieght threads managed by go runtime
import (
	"fmt"
	// "time"
)

func task(id int,w *sync.WaitGroup){

	defer w.Done()
	fmt.Println("Task",id,"is being processed")
}

func main(){
	var wg sync.WaitGroup
  for i :=0;i<=10;i++{
	// go task(i)

//  go	func (i int){
	wg.Add(1)
		go task(i,&wg)
	// }(i)
  }
 
wg.Wait()

}

