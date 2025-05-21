package main

import "fmt"

func main() {
	var value1 int32 = 32768
	var value2 int64 = int64(value1)
	// Number overflow
	var value3 int16 = int16(value1)

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
}