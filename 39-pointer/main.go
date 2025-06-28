package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by value
	var galih Address = Address{"Xuhui", "Shanghai", "China"}
	var vincent Address = galih // Copy value

	vincent.City = "Yangpu"
	
	fmt.Println(galih)
	fmt.Println(vincent)

	// Pointer: pass by reference
	var moe Address = Address{"Nantan", "Kyoto", "Japan"}
	var taka *Address = &moe // Referenced
	
	taka.City = "Kameoka"
	fmt.Println(moe, &moe)
	fmt.Println(taka, &taka)
}