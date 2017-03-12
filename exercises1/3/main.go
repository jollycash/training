package main

import "fmt"

func main() {
	fmt.Print("Please enter a number: ")
	var smallnr int64
	var biggernr int64
	fmt.Scan(&smallnr)
	fmt.Print("Now please enter a number bigger than the first one: ")
	fmt.Scan(&biggernr)
	bbs := biggernr % smallnr
	fmt.Printf("The remainder is %v \n", bbs)
}
