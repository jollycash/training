package main

import "fmt"

func main() {

	fmt.Print("Please enter your name: ")

	var name func() int32
	{
		fmt.Scan()
		return
	}
	fmt.Printf("Hello, %v \n", name)

}
