package main

import "fmt"

func sumAll(numbers ...int16) int16 {
	var total int16 = 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	var total int16 = sumAll(10, 20, 30, 40, 50, 100)
	fmt.Println(total)

	var nums [3]int16 = [3]int16{10, 20, 30} // Array
	// nums := [...]int16{10, 20, 30} // Short version
	var totalNums int16 = sumAll(nums[:]...)
	fmt.Println(totalNums)
}