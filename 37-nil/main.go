package main

import "fmt"

/*
	nil can only be used w/ certain data types,
	such as interface, func, map, slice, pointer,
	and channel.
*/

// func ItCouldNotWorks(name string) string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var taka map[string]string = NewMap("Taka")
	if taka == nil {
		fmt.Println("Data is empty.")
	} else {
		fmt.Println(taka["name"])
	}

	var moe map[string]string = NewMap("")
	if moe == nil {
		fmt.Println("Data is empty.")
	} else {
		fmt.Println(moe["name"])
	}
	/*
		If u return a nil value w/ a map
		data type based, then u couldn't assign
		a value into it. If so, better return 
		an empty map instead of nil.
	*/
	// This occurs panic
	// moe["name"] = "Moe"
}