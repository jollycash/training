package main

import "fmt"

func primenumbers(x int) []int {
	var primelist []int
	for i := 2; len(primelist) <= x; i++ {
		var factorlist []int
		// the following IF statement improves performance drastically when entering large values
		if i > 7 && i%2 == 0 || i > 7 && i%3 == 0 || i > 7 && i%5 == 0 || i > 7 && i%7 == 0 {
			continue
		}
		for o := 1; o <= i/2; o++ {
			if i%o == 0 {
				factorlist = append(factorlist, o)
			}
		}
		if len(factorlist) == 1 {
			primelist = append(primelist, i)
		}
	}
	return primelist
}

func main() {
	var a int
	fmt.Print("This program shows the n-th prime number. n = ")
	fmt.Scan(&a)
	fmt.Println(primenumbers(a)[a-1])
}
