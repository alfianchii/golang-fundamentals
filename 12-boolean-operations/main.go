package main

import "fmt"

func main() {
	const passedValue = 80
	var score uint8 = 90
	var absence uint8 = 80

	var isScorePassed bool = score > passedValue
	var isAbsencePassed bool = absence > passedValue

	var passed bool = isScorePassed && isAbsencePassed

	fmt.Println(passed)
}