package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToJP(address *Address) {
	address.Country = "Japan"
}

func main() {
	var moe Address = Address{"Nantan", "Kyoto", ""}
	ChangeCountryToJP(&moe)
	
	fmt.Println(moe)
}