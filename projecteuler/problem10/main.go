package main

import "fmt"

func primenumbers(x int) []int {
	var primelist []int
Outerloop:
	for i := 2; i <= x; i++ {
		// var factorlist []int
		// the following IF statement improves performance drastically when entering large values
		// however I'm aware this is still a very inefficient method for generating primes
		// if i > 7 && i%2 == 0 || i > 7 && i%3 == 0 || i > 7 && i%5 == 0 || i > 7 && i%7 == 0 {
		// 	continue
		// }
		for o := 2; o <= i/2; o++ {
			if i%o == 0 {
				continue Outerloop
			}
		}
		// if len(factorlist) == 1 {
		primelist = append(primelist, i)
		fmt.Println(i)
		// fmt.Println(primelist)
		// }
	}
	return primelist
}

func main() {
	var summ int
	for _, x := range primenumbers(2000000) {
		summ += x
	}
	fmt.Println(summ)
}
