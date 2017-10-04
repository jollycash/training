package main

import "fmt"

func triang(x int) int {
	var tri int
	for i := 1; i <= x; i++ {
		tri += i
	}
	return tri
}

func factors(x int) []int {
	var factorlist []int
	for i := 1; i <= x/2; i++ {
		if x%i == 0 && i != x/2 {
			factorlist = append(factorlist, i)
		} else if i == x/2 {
			factorlist = append(factorlist, i, x)
		}
	}

	return factorlist
}

func main() {
	for i := 100; true; i++ {
		if len(factors(triang(i))) == 500 {
			fmt.Println(triang(i))
			break
		}
	}
}
