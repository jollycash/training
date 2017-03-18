package main

import (
	"fmt"
)

func summultiples(z int) int {
	var sum int
	for i := 1; i <= z; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(summultiples(a))
}
