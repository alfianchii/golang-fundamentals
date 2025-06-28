package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var moe Address = Address{"Nantan", "Kyoto", "Japan"}
	var taka *Address = &moe
	fmt.Println(moe, taka)

	// var data1 *Address = &Address{}
	var data1 *Address = new(Address)
	var data2 *Address = data1
	fmt.Println(data1, data2)
}