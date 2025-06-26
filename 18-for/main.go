package main

import "fmt"

func main() {
	var counter int = 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// With init and post statements
	for counters := 1; counters <= 10; counters++ {
		fmt.Println("Perulangan ke", counter)
	}
}