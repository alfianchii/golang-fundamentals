package main

import "fmt"

func main() {
	var counter int = 0

	var increment func() = func() {
		fmt.Println("INCREMENTING!")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
}