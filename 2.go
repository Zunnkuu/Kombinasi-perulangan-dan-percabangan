package main

import (
	"fmt"
)

func main() {
	var input string
	// Membaca input dari pengguna
	fmt.Scanln(&input)

	// Menghitung jumlah vokal
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range input {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				break
			}
		}
		if char == '.' {
			break
		}
	}

	fmt.Print(count)
}
