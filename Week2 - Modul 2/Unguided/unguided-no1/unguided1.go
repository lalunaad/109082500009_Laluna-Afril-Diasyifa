package main 

import "fmt"

func main() {
	var tahun int

	fmt.Print("Tahun : ") 
	fmt.Scan(&tahun)

	var kabisat bool = false
	if(tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0) { 
		kabisat = true
    }
	fmt.Printf("Kabisat : %t\n", kabisat)
}
