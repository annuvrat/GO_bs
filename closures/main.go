package main



func counter()func()int{
	var count int = 0

	return func()int{
		count+= 1
		return count
	}
}

func main(){

	result:= counter()
	println(result()) //1
	println(result()) //2
	println(result()) //3
	  
}
