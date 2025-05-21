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
}