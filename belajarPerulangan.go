package main

import "fmt"

func main() {
	//perulangan dengan for
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	//perulanngan for dengan argumen kondisi
	var x = 0

	for x < 5 {
		fmt.Println("angka x: ", x)
		x++
	}

	//for dengan if
	var z = 0

	for {
		fmt.Println("angka z: ", z)
		z++
		if z == 5 {
			break
		}
	}

	//for dengan break dan continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	//label dalam perulangan
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
