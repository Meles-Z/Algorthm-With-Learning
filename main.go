package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// fileName:="test.txt"
	// file, err:=os.OpenFile(fileName, os.O_CREATE| os.O_WRONLY, 0644)
	// if err !=nil{
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()
	// b:=[]byte("Hello test of write to file")
	// n, err:=file.Write(b)
	// if err !=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Number of bytes written in file is: %d\n", n)
	// if err !=nil{
	// 	log.Fatal(err)
	// }

	// fileName := "test.txt"
	// file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// _, err=file.WriteString("Hello, Namaste of workld!")
	// if err !=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println("Done writting into file")

	// fileName:="test.txt"
	// data:="Hi everyone! What about your work!"
	//  err:=ioutil.WriteFile(fileName, []byte(data), 0 )
	//  if err !=nil{
	//     log.Fatal(err)
	//  }
	//  fmt.Println("Done!")
	// fileName:="test.txt"
	// file, err:=ioutil.ReadFile(fileName)
	// if err !=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(file))

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",content)

}
