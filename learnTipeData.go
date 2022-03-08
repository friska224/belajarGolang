package main

import "fmt"

func main() {
	var angkaPositif = 80
	var angkaNegatif = -90
	var desimal = 3.14
	var benarSalah bool = true
	var nama string = "friska"
	const name string = "Putri"

	fmt.Printf("output angka positif %d dan negatif %d\n", angkaPositif, angkaNegatif)
	fmt.Printf("output angka desimal %f\n", desimal)
	fmt.Printf("output angka desimal dengan 3 angka dibelakang koma %.3f\n", desimal)
	fmt.Printf("girls always? %t\n", benarSalah)
	fmt.Printf("nama : %s\n", nama)
	fmt.Print("print nama menggunakan cons : ", name, " gitu\n")

}
