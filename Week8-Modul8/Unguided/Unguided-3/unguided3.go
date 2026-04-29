package main

import "fmt"

func main() {
	var clubA, clubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&clubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&clubB)

	var results []string
	matchNum := 1

	for {
		var scoreA, scoreB int

		fmt.Printf("Pertandingan %d : ", matchNum)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			results = append(results, clubA)
		} else if scoreB > scoreA {
			results = append(results, clubB)
		} else {
			results = append(results, "Draw")
		}

		matchNum++
	}

	for i, result := range results {
		fmt.Printf("Hasil %d : %s\n", i+1, result)
	}

	fmt.Println("Pertandingan selesai")
}