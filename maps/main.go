package main

import (
	"fmt"
	
)

func main(){

	m:=make(map[string]string)
    n:=make(map[string]int)
	m["name"]="Annu"
	m["course"]="Golang"
    n["math"]=90
	fmt.Println(m["name"])
	fmt.Println(m["course"])
	// delete(m,"name")
	fmt.Println(m)
	clear(m)
	// clear(m)
	fmt.Println(n["mathw"])
} 