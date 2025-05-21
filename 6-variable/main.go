package main

import "fmt"

func main() {
	var name string

	name = "Alfian Taka"
	fmt.Println(name)
	
	name = "Taka"
	fmt.Println(name)

	// Auto detect data type
	var myName = "Alfian"
	fmt.Println(myName)
}