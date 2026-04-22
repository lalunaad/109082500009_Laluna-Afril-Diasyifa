package main

import "fmt"

type suhu float64

func CelciusToReamur(celcius suhu) suhu {
	return suhu(4.0 / 5.0 * float64(celcius))
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return suhu(9.0/5.0*float64(celcius) + 32)
}

func CelciusToKelvin(celcius suhu) suhu {
	return suhu(float64(celcius) + 273.15)
}

func main() {
	var masukan suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&masukan)

	hasilReamur := CelciusToReamur(masukan)
	hasilFahrenheit := CelciusToFahrenheit(masukan)
	hasilKelvin := CelciusToKelvin(masukan)

	fmt.Println()
	fmt.Printf("%g celcius = %g reamur\n", masukan, hasilReamur)
	fmt.Printf("%g celcius = %g fahrenheit\n", masukan, hasilFahrenheit)
	fmt.Printf("%g celcius = %g kelvin\n", masukan, hasilKelvin)
}