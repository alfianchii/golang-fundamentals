package main

import "fmt"

func Ups() interface{} {
	return []int{1, 2, 3} // Slice
}

func Woops() any {
	return "'any' means empty interface" // String
}

func main() {
	var empty any = Ups()
	fmt.Println(empty)

	var emptyAgain interface{} = Woops()
	fmt.Println(emptyAgain)
}