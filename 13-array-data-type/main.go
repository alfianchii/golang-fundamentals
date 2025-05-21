package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Alfian"
	names[1] = "Taka"
	
	fmt.Println("First name:", names[0])
	fmt.Println("Second name:", names[1])

	// Got error
	// names[2] = "T4mv4n"
	
	names[1] = "Chii"
	fmt.Println("Second name:", names[1])
}