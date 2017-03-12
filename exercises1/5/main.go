package main

import "fmt"

func main() {
	for i := 101; i > 2; i-- {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%v \n", i)
		}
	}
}
