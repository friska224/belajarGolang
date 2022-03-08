package main

import "fmt"

func main() {
	//cara menggunakan map
	var ayam map[string]int
	ayam = map[string]int{}

	ayam["januari"] = 50
	ayam["februari"] = 40

	fmt.Println("januari", ayam["januari"])
	fmt.Println("mei", ayam["mei"])

	//inisialisasi
	var data map[string]int
	data["satu"] = 1

	//inisialisasi yg salah
	/*data = map[string]int{}
	data["satu"] = 1*/

	// cara inisialisasi horizontal
	var ayam1 = map[string]int{"januari": 50, "februari": 40}

	// cara inisialisasi vertical
	var ayam2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	//inisialisasi tanpa nilai awal
	var ayam3 = map[string]int{}
	var ayam4 = make(map[string]int)
	var ayam5 = *new(map[string]int)

	//penggunakan for range di map
	for key, val := range chicken2 {
		fmt.Println(key, "  \t:", val)
	}

	//menghapus item map

	fmt.Println(len(ayam)) // output 2
	fmt.Println(ayam)

	delete(ayam, "januari") //hapus disini

	fmt.Println(len(ayam)) // cek ayam tinggal 1
	fmt.Println(ayam)

	//deteksi item dengan key tertentu
	var value, isExist = ayam["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	//kombinasi map dan slice
	var ayams = []map[string]string{
		map[string]string{"nama": "ayam biru", "gender": "jantan"},
		map[string]string{"name": "ayam merah", "gender": "jantan"},
		map[string]string{"nama": "ayam kuning", "gender": "betina"},
	}

	for _, ayam := range ayams {
		fmt.Println(ayam["gender"], ayam["nama"])
	}

}
