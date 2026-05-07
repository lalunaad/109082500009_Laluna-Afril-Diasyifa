package main

import "fmt"

type arrReal [1000]float64

func terkecil(tabReal arrReal, n int) float64 {
	var min float64 = tabReal[0]
	var j int = 1

	for j < n {
		if min > tabReal[j] {
			min = tabReal[j]
		}
		j = j + 1
	}

	return min
}

func terbesar(tabReal arrReal, n int) float64 {
	var max float64 = tabReal[0]
	var j int = 1

	for j < n {
		if max < tabReal[j] {
			max = tabReal[j]
		}
		j = j + 1
	}

	return max
}

func main() {
	var n int
	var berat arrReal

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := terkecil(berat, n)
	max := terbesar(berat, n)

	fmt.Printf("%.2f %.2f\n", min, max)
}