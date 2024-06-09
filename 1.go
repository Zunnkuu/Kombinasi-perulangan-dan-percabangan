package main

import (
	"fmt"
)

func main() {
	var sum, first, second int
	for {
		fmt.Scan(&first, &second)
		if first == 0 && second == 0 {
			break
		}
		if first%2 != 0 && second%2 == 0 {
			sum += first + second
		}
	}

	fmt.Println(sum)
}
