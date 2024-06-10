package main

import (
	"fmt"
	"math"
)

func isPrimeZ(n int) bool {
	factors := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				factors++
			} else {
				factors += 2
			}
		}
	}
	return factors%2 == 1
}

func main() {
	var num int

	fmt.Scan(&num)
	fmt.Println(isPrimeZ(num))
}
