package main

import (
	"fmt"
	"math"
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

func summit(z ...int) int {
	var summ int
	for _, i := range z {
		summ += i
	}
	return summ
}

func main() {
	fmt.Println(summit(converse(strconv.FormatFloat(math.Pow(2, 1000), 'f', 1, 64))...))
}
