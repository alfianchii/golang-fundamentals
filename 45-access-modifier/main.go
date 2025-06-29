package main

import (
	"45-access-modifier/helper"
	"fmt"
)

func main() {
	var result string = helper.SayHello("Taka")

	fmt.Println(result)
	fmt.Println(helper.App)
	// // Error occurs
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye("Galih"))
	helper.Example()
}