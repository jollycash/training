package main

import "fmt"

func main() {
	v := (true && false) || (false && true)
	fmt.Printf("%v", v)
}
