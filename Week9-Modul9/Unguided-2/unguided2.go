package main

import "fmt"

type arrReal [1000]float64

func hitungTotalWadah(arr arrReal, x, y int) (arrReal, int) {
	var totalWadah arrReal
	var i, j, idx int

	idx = 0

	for i = 0; i < x; i = i + y {
		var total float64 = 0

		for j = 0; j < y; j++ {
			if i+j < x {
				total = total + arr[i+j]
			}
		}

		totalWadah[idx] = total
		idx = idx + 1
	}

	return totalWadah, idx
}

func hitungRataRata(totalWadah arrReal, numWadah int) float64 {
	var sum float64 = 0
	var i int

	for i = 0; i < numWadah; i++ {
		sum = sum + totalWadah[i]
	}

	return sum / float64(numWadah)
}

func main() {
	var x, y int
	var berat arrReal

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	totalWadah, numWadah := hitungTotalWadah(berat, x, y)

	for i := 0; i < numWadah; i++ {
		fmt.Printf("%.2f", totalWadah[i])

		if i < numWadah-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println()

	rataRata := hitungRataRata(totalWadah, numWadah)
	fmt.Printf("%.2f\n", rataRata)
}