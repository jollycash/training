package main

import "fmt"

// this function sums all even numbers in a list of integers
func summit(z ...int) int {
	sumeven := 0
	for _, v := range z {
		if v%2 == 0 {
			sumeven = sumeven + v
		}
	}
	return sumeven
}

func main() {
	nrold := 0
	nradd := 1
	var nrlist []int
	// the function in the range argument simulates the Fibonacci sequence
	for i := 1; i < 40000001; func() {
		nradd = nrold + i
		nrold = i
		i = nradd
	}() {
		nrlist = append(nrlist, nradd) //alternatively could already filter for even numbers here with IF statement
	}
	fmt.Println("All Fibonacci numbers between 1 and 4mio:", nrlist)
	fmt.Println("Sum of all even Fibonacci numbers between 1 and 4mio:", summit(nrlist...))
}
