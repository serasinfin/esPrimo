package main

import (
	"fmt"
	"math"
)

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Introduce un número: ")
	fmt.Scan(&num)

	for i := 2; i <= num; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
}
