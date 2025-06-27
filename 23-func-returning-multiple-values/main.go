package main

import "fmt"

func getMyFullName() (string, string) {
	return "Alfian", "Taka"
}

func main() {
	var myFirstName, myLastName string = getMyFullName()
	fmt.Println(myFirstName, myLastName)
}