package main

import "fmt"

func sumsq(z int) int {
	var sumnr int
	for i := 1; i <= z; i++ {
		sumnr += i * i
	}
	return sumnr
}

func sqsum(z int) int {
	var sumnr int
	for i := 1; i <= z; i++ {
		sumnr += i
	}
	return sumnr * sumnr
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%v - %v = %v \n", sumsq(a), sqsum(a), sqsum(a)-sumsq(a))
}
