package main

import "fmt"

func greatestvalue(z ...int) int {
	comparison := 0
	for _, v := range z {
		if v > comparison {
			comparison = v
		}
	}
	return comparison
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println("The biggest number is: ", greatestvalue(a, b, c, d))
}
