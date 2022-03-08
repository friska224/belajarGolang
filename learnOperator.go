package main

import "fmt"

func main() {
	var hasil = 4 - 2
	var isEqual = (hasil == 2)
	var value = (((2+6)%3)*4 - 2) / 3

	fmt.Printf("nilai %d (%t) || %d\n", hasil, isEqual, value)

	var kanan = true
	var kiri = false

	var kananDanKiri = kanan && kiri
	var kananAtauKiri = kanan || kiri
	var bukanKanan = !kanan

	fmt.Printf("kanan && kiri \t (%t)\n", kananDanKiri)
	fmt.Printf("kanan || kiri \t (%t)\n", kananAtauKiri)
	fmt.Printf("bukan kanan \t (%t)\n", bukanKanan)
}
