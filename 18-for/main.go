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

	// Range (collection date types: array, slice, map)
	var names []string = []string{"Taka", "Ogi", "Moe"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for index, value := range names {
		fmt.Println(index, value)
	}
	for _, value := range names {
		fmt.Println(value)
	}
}