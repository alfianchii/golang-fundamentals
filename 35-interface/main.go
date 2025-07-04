package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHelloTo(val HasName) {
	fmt.Println("Hello, " + val.GetName() + "!")
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var moe Person = Person{Name: "Moe"}
	SayHelloTo(moe)

	var cat Animal = Animal{Name: "Cat"}
	SayHelloTo(cat)
}