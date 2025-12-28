package main

import "fmt"

func main(){

	m:=make(map[string]string)

	m["name"]="Annu"
	m["course"]="Golang"

	fmt.Println(m["name"])
	fmt.Println(m["course"])
	// clear(m)
	fmt.Println(m)
}