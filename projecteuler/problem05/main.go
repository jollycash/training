package main

import "fmt"

// func main() {
// 	var i int
// 	for {
// 		i++
// 		if i%1 == 0 && i%2 == 0 && i%3 == 0 && i%4 == 0 && i%5 == 0 && i%6 == 0 && i%7 == 0 && i%8 == 0 && i%9 == 0 && i%10 == 0 && i%11 == 0 && i%12 == 0 && i%13 == 0 && i%14 == 0 && i%15 == 0 && i%16 == 0 && i%17 == 0 && i%18 == 0 && i%19 == 0 && i%20 == 0 {
// 			break
// 		}
// 	}
//
// 	fmt.Println(i)
// }

func main() {
	var a int
	var b int
	fmt.Scan(&a)
Outerloop:
	for i := 1; i > 0; i++ {
		for j := 1; j <= a; j++ {
			if i%j != 0 {
				continue Outerloop
			}
		}
		b = i
		break
	}

	fmt.Println(b)
}