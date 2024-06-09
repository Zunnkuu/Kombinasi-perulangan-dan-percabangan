package main

import (
	"fmt"
)

func main() {
	var episode, duration, totalDuration int
	for {
		// Membaca input nomor episode dan durasi episode
		fmt.Scan(&episode, &duration)

		// Jika pasangan yang dibaca adalah 0 0, hentikan pembacaan
		if episode == 0 && duration == 0 {
			break
		}

		// Jika nomor episode ganjil, tambahkan durasinya ke total
		if episode%2 != 0 {
			totalDuration += duration
		}
	}

	fmt.Print(totalDuration)
}
