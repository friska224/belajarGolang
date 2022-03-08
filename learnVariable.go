package main

import "fmt"

func main() {
	var namaDepan string = "Friska"
	var namaBelakang string
	namaBelakang = "Putri"

	namaTengah := "Sri Eka"

	var satu, dua, tiga string

	satu, dua, tiga = "siji", "loro", "telu"

	nama := new(string)

	fmt.Printf("Hallooo %s %s!!\n", namaDepan, namaBelakang)
	fmt.Println("Hello pakai Println ", namaDepan, namaBelakang+"!!")
	fmt.Printf("Hellow pakai deklarasi variable tanpa tipe data %s %s %s\n", namaDepan, namaTengah, namaBelakang)
	fmt.Printf("Ini pakai multi variable : %s %s %s \n", satu, dua, tiga)
	fmt.Println(*nama)
}
