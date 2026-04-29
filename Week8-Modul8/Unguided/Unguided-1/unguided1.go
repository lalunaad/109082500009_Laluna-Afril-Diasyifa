package main

import (
	"fmt"
	"math"
)

type Titik struct {
	koordinat [2]int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	dx := float64(p.koordinat[0] - q.koordinat[0])
	dy := float64(p.koordinat[1] - q.koordinat[1])

	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var lingkaran1 Lingkaran
	var lingkaran2 Lingkaran
	var titik Titik

	fmt.Scan(&lingkaran1.pusat.koordinat[0], &lingkaran1.pusat.koordinat[1], &lingkaran1.r)
	fmt.Scan(&lingkaran2.pusat.koordinat[0], &lingkaran2.pusat.koordinat[1], &lingkaran2.r)
	fmt.Scan(&titik.koordinat[0], &titik.koordinat[1])

	dalam1 := didalam(lingkaran1, titik)
	dalam2 := didalam(lingkaran2, titik)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}