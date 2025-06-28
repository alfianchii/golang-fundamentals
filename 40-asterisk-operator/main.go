package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pointer: pass by reference
	var moe Address = Address{"Nantan", "Kyoto", "Japan"}
	var taka *Address = &moe // Referenced
	
	taka.City = "Kameoka"
	fmt.Println("REAL:", moe, "POINTER:", &moe)
	fmt.Println("REAL:", taka, "POINTER:", &taka, "DEREFERENCE:", *taka)

	// x means access the actual value
	// &x means access the address of x
	// *x means access the 'actual value' it points to
	*taka = Address{"Fukuchiyama", "Kyoto", "Japan"} // Therefore it would change the moe's values

	fmt.Println("REAL:", moe, "POINTER:", &moe)
	fmt.Println("REAL:", taka, "POINTER:", &taka, "DEREFERENCE:", *taka)
}