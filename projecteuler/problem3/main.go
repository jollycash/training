package main

import "fmt"

func primenumbers(x int) []int {
	var primelist []int
	for i := 2; i <= x; i++ {
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

// func primefactors(z int) []int {
// 	number := z
// 	var primefactorlist []int
// 	for i := z; i != 0; func() {
// 		for _, v := range primenumbers(z) {
// 			if number%v == 0 && number/v != 1 {
// 				primefactorlist = append(primefactorlist, v)
// 				number = number / v
// 			} else if number%v == 0 && number/v == 1 {
// 				primefactorlist = append(primefactorlist, v)
// 				i = number
// 				break
// 			}
// 		}
// 	}() {
// 	}
// 	return primefactorlist
// }

func primefactors(z int) []int {
	number := z
	var primefactorlist []int
Outerloop:
	for {
		if number == 0 {
			break
		} else {
			for _, v := range primenumbers(z) {
				// for _, v := range primenumbers(z) {
				if number%v == 0 {
					primefactorlist = append(primefactorlist, v)
					number = number / v
					if number == 1 {
						break Outerloop
					}
					continue Outerloop
				}
			}
		}
	}
	return primefactorlist
}

func largestvalue(z ...int) int {
	comparison := 0
	for _, v := range z {
		if v > comparison {
			comparison = v
		}
	}
	return comparison
}

func main() {
	var a int
	fmt.Print("Enter any natural number above 2: ")
	fmt.Scan(&a)
	fmt.Println("The prime factors are: ", primefactors(a))
	fmt.Println("The largest prime factor is:", largestvalue(primefactors(a)...))
}
