package main

import "fmt"

type Person struct {
	Name string
}

func (data *Person) Married() {
	data.Name = "Mr. " + data.Name
}

func main() {
	var galih Person = Person{"Galih"}
	fmt.Println(galih)

	// (&galih).Married()
	galih.Married()

	fmt.Println(galih)
}