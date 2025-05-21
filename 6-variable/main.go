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

	// Replace 'var' w/ ':=' keyword
	title := "君の名前は？"
	fmt.Println(title)

	title = "よつばと！"
	fmt.Println(title)

	// Multiple vars
	var (
		firstName = "Alfian"
		lastName = "Taka"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}