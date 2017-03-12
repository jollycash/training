package main

import "fmt"

func main() {
	x := func(z int) (int, bool) {
		zidivided := z / 2
		zeven := z%2 == 0
		return zidivided, zeven
	}
	var a int
	fmt.Scan(&a)
	fmt.Println(x(a))
}
