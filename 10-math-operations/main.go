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

	a += 904
	fmt.Println("The value of 'a' rn:", a)

	a -= 120
	fmt.Println("The value of 'a' rn:", a)

	a *= 10
	fmt.Println("The value of 'a' rn:", a)

	a -= 2
	fmt.Println("The value of 'a' rn:", a)

	a %= 2
	fmt.Println("The value of 'a' rn:", a)
}