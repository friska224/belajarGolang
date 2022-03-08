package main

import (
	"fmt"
	"math"
)

func main() {
	//inisialisasi variable
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	//pemanggilan fungsi
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

//fungsi kalkulator
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}
