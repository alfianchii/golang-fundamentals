package main

import "fmt"

func thisReturnsString() interface{} {
	return "OK"
}

func thisReturnsInt() interface{} {
	return 10
}

func main() {
	var result interface{} = thisReturnsString()
	var resultInString string = result.(string)
	fmt.Println(resultInString)

	var result2 interface{} = thisReturnsInt()
	var resultInInt int = result2.(int)
	fmt.Println(resultInInt)

	// // Panic occurs
	// var data int = result.(int)
	// fmt.Println(data)

	// Type Assertions w/ switch statement
	var data interface{} = thisReturnsInt()
	switch value := data.(type) {
	case string:
		fmt.Println("This is a 'string' data type:", value)
	case int:
		fmt.Println("This is a 'int' data type:", value)
	default:
		fmt.Println("Unknown:", value)
	}
}