package main

import (
	"fmt"
	"maps"
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
	k:= map[string]int{"price":100}
    k2:= map[string]int{"price":100}


	fmt.Println(maps.Equal(k,k2))



	v,ok:= k["price"]


	fmt.Println(v)
	if ok{
		fmt.Println("present")
	}else{
		fmt.Println("not present")
	}
	fmt.Println(k["price"])


} 