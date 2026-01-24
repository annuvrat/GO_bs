package main

import "fmt"

// import "fmt"

// // func add(a, b int)int{
// // 	return a+b
// // }

// func main(){
// 	// result:= add(3,4)
// lang1,lang2,lang3:= lang()
// fmt.Println(lang1)
// fmt.Println(lang2)
// fmt.Println(lang3)

// 	// fmt.Println("The sum is:",result)
// 	fmt.Println(lang())
// }

// func lang()(string,string,string){
// 	return "GoLanguage","Golang","Go"
// }


func closure()func(a int)int{
	return func(a int)int{
		return a*a
	}
}


func main(){
  
	result:= closure()
	fmt.Println(result(3))

}

