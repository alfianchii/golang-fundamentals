package main

import "fmt"

func main() {
	const firstName string = "Alfian"
	const lastName = "Taka"
	
	fmt.Println(firstName)
	fmt.Println(lastName)

	// Error: immutable
	// firstName = "Alfian Tampan"
	// lastName = "Bamnget"

	// Multiple const
	const (
		str1 = "最近、元気？"
		str2 = "うん"
	)

	fmt.Println(str1)
	fmt.Println(str2)
}