package main

import "fmt"

func main() {
	var counter int = 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
}