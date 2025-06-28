package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHelloTo(name string) {
	fmt.Println("Hello, " + name + ". My name is " + customer.Name + "~")
}

func main() {
	var taka Customer = Customer{Name: "Taka", Address: "Indonesia", Age: 30}
	taka.sayHelloTo("Michael")
}