package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	var j int = 1

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for j < n {
		if *bMin > arrBerat[j] {
			*bMin = arrBerat[j]
		}

		if *bMax < arrBerat[j] {
			*bMax = arrBerat[j]
		}

		j = j + 1
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var sum float64 = 0
	var i int = 0

	for i < n {
		sum = sum + arrBerat[i]
		i = i + 1
	}

	return sum / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d : ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)
	avg := rerata(berat, n)

	fmt.Printf("\nBerat balita minimum : %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", max)
	fmt.Printf("Rerata berat balita : %.2f kg\n", avg)
}