package main

import (
	"fmt"
	"strconv"
	"strings"
)

func converse(x string) []int {
	splitstrings := strings.Split(x, "")
	var splitnumbers []int
	for _, v := range splitstrings {
		if i, err := strconv.Atoi(v); err == nil {
			splitnumbers = append(splitnumbers, i)
		}
	}
	return splitnumbers
}

func largestvalue(y int, z ...int) int {
	largest := 0
	for i, j := 0, len(z); i <= j-y; i++ {
		currentsliceprod := 1
		for k := i; k <= i+y-1; k++ {
			currentsliceprod = currentsliceprod * z[k]
		}
		if currentsliceprod > largest {
			largest = currentsliceprod
		}
	}
	return largest
}

func main() {
	var a string
	var b int
	fmt.Scan(&a)
	fmt.Println(converse(a))
	fmt.Scan(&b)
	fmt.Println(largestvalue(b, converse(a)...))
}
