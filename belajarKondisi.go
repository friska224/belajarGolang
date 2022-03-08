package main

import (
	"fmt"
)

func main() {

	//kondisi If Else
	var nilai = 10

	if nilai == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if nilai > 5 {
		fmt.Println("lulus")
	} else if nilai == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", nilai)
	}

	//variable temp pada if else
	var point = 9870.0

	if persen := point / 100; persen >= 100 {
		fmt.Printf("%.1f erfect\n", persen)
	} else if persen >= 70 {
		fmt.Printf("%.1f good\n", persen)
	} else {
		fmt.Printf("%.1f not bad\n", persen)
	}

	//kondisi dengan switch case
	var isi = 7

	switch isi {
	case 8:
		fmt.Println("good\n")
	case 9:
		fmt.Println("awsome\n")
	case 10:
		fmt.Println("perfect\n")
	default:
		fmt.Println("not bad\n")
	}

	//switch case dengan banyak kondisi
	var val = 7

	switch val {
	case 5, 6, 7, 8:
		fmt.Println("good\n")
	case 10:
		fmt.Println("perfect\n")
	default:
		fmt.Println("not bad\n")
	}

	//switch dengan gaya if else
	switch {
	case val == 8:
		fmt.Println("perfect")
	case (val < 8) && (val > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//fallthough dalam switch
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//seleksi kondisi
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
