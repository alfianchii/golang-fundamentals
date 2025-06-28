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

	// Panic occurs
	var data int = result.(int)
	fmt.Println(data)
}