package main

import "fmt"

func factorialLoop(val int) int {
	var result int = 1

	for i := val; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(val int) int {
	if (val == 1) {
		return 1
	} else {
		return val * factorialRecursive(val - 1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}