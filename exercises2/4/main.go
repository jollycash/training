package main

import "fmt"

func foo(z ...int) {
	comparison := 0
	for _, v := range z {
		if v > comparison {
			comparison = v
		}
	}
	fmt.Println(comparison)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
