package main

import "fmt"

func main() {
	//slice
	var buah = []string{"apel", "anggur", "pisang", "melon"}
	fmt.Println(buah[0])

	var buahBaru = buah[0:2]

	fmt.Println(buahBaru)

	//slice = tipe data referance
	var aBuah = buah[0:3]
	var bBuah = buah[1:4]

	var aaBuah = aBuah[1:2]
	var baBuah = bBuah[0:1]

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(baBuah)

	//fungsi len untuk melihat jumlah elemen
	fmt.Println(len(buah))
	fmt.Println(len(buahBaru))

	//cap untuk melihat jumlah maksimum dari slice
	fmt.Println(cap(buah))
	fmt.Println(cap(buahBaru))

	//fungsi append
	var cBuah = append(buah, "papaya")

	fmt.Println(buah)
	fmt.Println(cBuah)

	//fungsi copy
	contohMake := make([]string, 3)
	kopiannya := []string{"aa", "bb", "cc", "dd"}
	n := copy(contohMake, kopiannya)

	fmt.Println(contohMake)
	fmt.Println(kopiannya)
	fmt.Println(n)

	//akses slice dengan 3 elemen
	var adaLagi = buah[0:2:2]
	fmt.Println(adaLagi)
	fmt.Println(len(adaLagi))
	fmt.Println(cap(adaLagi))
}
