package main

import "fmt"

func main() {
	// Make a new 'data type'
	type NoKTP string

	var ktpTaka NoKTP = "90148714900"
	var ktpKanji string = "139913091"
	var convertKTPKanji NoKTP = NoKTP(ktpKanji)

	fmt.Println(ktpTaka)
	fmt.Println(ktpKanji)
	fmt.Println(convertKTPKanji)
}