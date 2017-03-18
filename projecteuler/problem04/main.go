package main

import (
	"fmt"
	"math"
)

// this function reverses an int
// e.g. in = "123" out = "321"
func reverse(z int) int {
	var a int
	var c int
	var d int
	for {
		if z < 10 {
			d = d + z
			break
		}
		a = z % 10
		z = z / 10
		c = a + d
		d = c * 10
	}
	return d
}

// this function looks for the largest product of all
// combinations of two n-digit numbers, where n=z

func palinprod(z int) int {
	var palinprod int
Outerloop:
	for i := z - 1; i > 0; i-- {
		for k := i; k > 1; k-- {
			if i*k == reverse(i*k) {
				if i*k > palinprod {
					palinprod = i * k
				} else if i*k < palinprod {
					continue Outerloop
				}
			}
		}
	}
	return palinprod
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(reverse(a))
	fmt.Println(a == reverse(a))
	fmt.Println(math.Pow10(3))
	fmt.Println(palinprod(int(math.Pow10(a))))
}
