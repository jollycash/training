package main

import "fmt"

func summit(y, z *int) {
	*y = *y + *z
	return
}

func main() {
	x := 0
	for i := 0; i < 1001; i++ {
		switch i % 3 {
		case 0:
			summit(&x, &i)

		}
		switch i % 5 {
		case 0:
			summit(&x, &i)
		}
	}
	fmt.Printf("%v \n", x)
}
