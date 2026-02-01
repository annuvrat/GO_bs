package main

import (
	// "fmt"
	"fmt"
	"os"
)

func main(){

// f,err := os.Open("t.txt")

// if err!=nil{
// 	//log the error
// 	panic(err)
// }

//   fileInfo ,err:=f.Stat()

//   if err!=nil{
// 	panic(err)
//   }

//   fmt.Println("file name",fileInfo.Name())
//   fmt.Println("file or folder",fileInfo.IsDir())
//   fmt.Println("file size:",fileInfo.Size())
//     fmt.Println("perms",fileInfo.Mode())
// 	  fmt.Println("file mod at",fileInfo.ModTime())

//read file
// f,err:=  os.Open("t.txt")
// if err!=nil{
// 	panic(err)	


// }

// defer f.Close()

// buf:= make([]byte,13)

// d,err :=  f.Read(buf)

// if err!= nil{
// 	panic(err)
// }

// for i:=0;i<len(buf);i++{

// println("data",d,string(buf[i]))

// }

// more direct approach not adivisable always only for small files
// f,err := os.ReadFile("t.txt")

// if err!=nil{
// 	panic(err)
// }
// fmt.Println("data",string(f))

//read folders
dir,err:= os.Open("./")

if err != nil{
	panic(err)
}

defer dir.Close()

fileINfo,err:= dir.ReadDir(2)

for _,file:= range fileINfo {
	fmt.Println("name",file.Name())
}

}