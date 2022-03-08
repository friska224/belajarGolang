package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	//IIFE
	var nilai = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var nilaiBaru = func(min int) []int {
		var r []int
		for _, e := range nilai {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", nilai)
	fmt.Println("filtered number :", nilaiBaru)

	//memanggil closure return
	var mak = 3
	var angka = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var jml, getNumbers = findMax(angka, mak)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", jml)
	fmt.Println("value \t:", theNumbers)
}

//closure sebagai return
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
