package main

import "fmt"

func sumeven(z ...int) int {
	sumeven := 0
	for _, v := range z {
		if v%2 == 0 {
			sumeven = sumeven + v
		}
	}
	return sumeven
}

func main() {
	var a int
	var b int
	c := 1
	var nrlist []int
	for {
		fmt.Println(nrlist)
		nrlist = append(nrlist, c)
		a = b
		b = c
		c = a + b
		if c > 4000000 {
			break
		}
	}
	fmt.Println("Sum of all even Fibonacci numbers between 1 and 4mio:", sumeven(nrlist...))
}
