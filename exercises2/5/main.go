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
	nro := 0
	nr := 1
	var nrlist []int
	// nrlist := []int{1}
	for i := 1; i < 40000001; func() {
		nr = nro + i
		nro = i
		i = nr
	}() {
		nrlist = append(nrlist, nr)
	}
	fmt.Println("These are all Fibonacci numbers between 1 and 4mio:", nrlist)
	fmt.Println("This is the sum of all even Fibonacci numbers between 1 and 4mio:", summit(nrlist...))
}
