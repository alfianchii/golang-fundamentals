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
	fmt.Println("--------")

	// Immediately state the value of array
	// var myName [2]string = [2]string{
	// 	"Alfian",
	// 	"Taka",
	// }
	var myName = [3]string{"Alfian", "Taka"}
	fmt.Println(myName)
	fmt.Println("--------")

	// Array funcs
	// [...] calculate based on the early values
	var yourName = [...]string{"Earl", "Phantomhive"}
	var lengthOfYourName int = len(yourName)
	// Got error
	// yourName[2] = "Freak"

	fmt.Println(lengthOfYourName)
}