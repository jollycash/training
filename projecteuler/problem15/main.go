package main

import (
	"fmt"
	"strconv"
)

func multip(x, y float64) float64 {
	newx := x + y
	newy := x
	for i := 1; i < int(x); i++ {
		newx = newx * ((x + y) - float64(i))
		fmt.Println("1", newx)
	}
	for i := x - 1; i > 1; i-- {
		newy = newy * i
		fmt.Println("2", newy)
	}

	return newx / newy
}

func bigger(x, y float64) (float64, float64) {
	var bigger float64
	var smaller float64
	if x > y {
		bigger = x
		smaller = y
	} else {
		bigger = y
		smaller = x
	}
	return bigger, smaller
}

func main() {
	var a float64
	var b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(strconv.FormatFloat(multip(bigger(a, b)), 'f', 1, 64))
}
