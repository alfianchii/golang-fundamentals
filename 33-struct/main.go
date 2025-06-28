package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var taka Customer

	fmt.Println(taka)

	taka.Name = "Alfian Taka"
	taka.Address = "Indonesia"
	taka.Age = 69

	fmt.Println(taka)
	fmt.Println(taka.Name)
	fmt.Println(taka.Address)
	fmt.Println(taka.Age)
}