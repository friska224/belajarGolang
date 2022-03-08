package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//penerapan fungsi printMassage
	var names = []string{"Friska", "Putri"}
	printMessage("halo", names)

	//penerapan fungsi dengan nilai balik
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
}

//fungsi baru bernama printMassage
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

//fungsi randomWithRange dengan nilai balik (return)
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
