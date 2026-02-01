package main

import (
	"fmt"
	"os"
)

// "bufio"
// "fmt"
// "os"

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
// dir,err:= os.Open(".")

// if err != nil{
// 	panic(err)
// }

// defer dir.Close()

// fileINfo,err:= dir.ReadDir(2)

// for _,file:= range fileINfo {
// 	fmt.Println("name",file.Name(),file.IsDir())
// }


//create the file and folders


// file,err := os.Create("bitch.txt")

// if err!= nil{

// panic(err)
// }
// defer file.Close()

// // file.WriteString("hi biches")
// // file.WriteString("bitch lasagnia")

// bytes := []byte("hello go")

// file.Write(bytes)

//read and write to another file( streaming fashion)

// file,err:= os.Open("t.txt")


// if err!= nil{

// 	panic(err)

// }

// defer  file.Close()

// destFile ,err := os.Create("t1.txt")

// if err!=nil {
// 	panic(err)
// }

// defer destFile.Close()

// reader := bufio.NewReader(file)
// writer:= bufio.NewWriter(destFile)

// for{
// 	b,err:=reader.ReadByte()

// 	if err!=nil{

// 		if err.Error()!= "EOF"{

// 			panic(err)
// 		}
// 		break
// 	}


// 	er:=writer.WriteByte(b)

// 	if er!= nil{
// 		panic(er)
// 	}



// }

// writer.Flush()

// fmt.Println("written to new file")


//// delete a file


// file,err:= os.Open("t.txt")


// if err!= nil{

// 	panic(err)

// }

// defer  file.Close()

err:= os.Remove("t1.txt")

 if err!= nil{
	panic(err)
 }


 fmt.Println("file deleted")

}