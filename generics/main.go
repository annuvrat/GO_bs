package main



func printslice[T any](items []T) {

	for _, item := range items {
		println(item)
	}

}

func main() {

	// nums:=[]int{1,2,3,4,5}

	names:= []string{"Alice", "Bob", "Charlie"}

	printslice(names)
  
	
}
