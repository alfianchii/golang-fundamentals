package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := a + b
	fmt.Println("The result of c:", c)

	d := 20
	e := 103
	f := 90
	g := d - e * f / a
	fmt.Println("The result of g:", g)	
}