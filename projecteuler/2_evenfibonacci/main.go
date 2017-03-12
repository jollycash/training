package main

import "fmt"

func summit(z ...int) int {
	greatest := 0
	for _, v := range z {
		if v%2 == 0 {
			greatest = greatest + v
		}
	}
	return greatest
}

func main() {
	nrold := 0
	nradd := 1
	var nrlist []int
	// nrlist := []int{1}
	for i := 1; i < 40000001; func() {
		nradd = nrold + i
		nrold = i
		i = nradd
	}() {
		nrlist = append(nrlist, nradd)
	}
	fmt.Println("All Fibonacci numbers between 1 and 4mio:", nrlist)
	fmt.Println("Sum of all even Fibonacci numbers between 1 and 4mio:", summit(nrlist...))
}
