package main

import (
	"fmt"
)

func main() {
	var number int
	// Membaca input dari pengguna
	fmt.Scan(&number)
	// Menentukan apakah bilangan tersebut adalah bilangan prima
	isPrime := true

	if number <= 1 {
		isPrime = false
	} else {
		if number == 2 {
			isPrime = true
		} else if number%2 == 0 {
			isPrime = false
		} else {
			for i := 3; i*i <= number; i += 2 {
				if number%i == 0 {
					isPrime = false
					break
				}
			}
		}
	}

	fmt.Println(isPrime)
}
