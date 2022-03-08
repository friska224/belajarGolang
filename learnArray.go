package main

import "fmt"

func main() {
	//array
	var names [4]string
	names[0] = "friska"
	names[1] = "sri"
	names[2] = "eka"
	names[3] = "putri"

	fmt.Println(names[0], names[1], names[2], names[3])

	//inisialisiasi di awal
	var buah = [4]string{"apel", "anggur", "pisang", "melon"}

	fmt.Println("Jumlah element \t\t", len(buah))
	fmt.Println("Isi semua element \t", buah)

	//inisialisasi vertikal
	var bunga = [4]string{
		"mawar",
		"melati",
		"kamboja",
		"tulip",
	}
	fmt.Println("Isi semua element bunga \t", bunga)

	//inisialisasi tanpa jumlah elemen
	var angka = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", angka)
	fmt.Println("jumlah elemen \t:", len(angka))

	//array multidimensi
	var angka1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var angka2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("nomor1", angka1)
	fmt.Println("nomor2", angka2)

	//perulangan elemen array
	var bebuahan = [4]string{"jeruk", "strawberry", "durian", "mangga"}

	for i := 0; i < len(bebuahan); i++ {
		fmt.Printf("array bebuahan %d : %s\n", i, bebuahan[i])
	}

	//perulangan for range pada array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	//menggunakan make
	var minuman = make([]string, 2)
	minuman[0] = "teh"
	minuman[1] = "kopi"

	fmt.Println(minuman)

}
