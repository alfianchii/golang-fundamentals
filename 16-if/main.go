package main

import "fmt"

func main() {
	var name string = "Taka"

	if name == "Alfian" {
		fmt.Println("Halo, " + name)
	} else if name == "Taka" {
		fmt.Println("こんにちは, " + name)
	} else if name == "Ogi" {
		fmt.Println("Hello, " + name)
	} else if name == "Taufan" {
		fmt.Println("สวัสดี, " + name)
	} else {
		fmt.Println("Hello! What's ur name?")
	}

	if nameLength := len(name); nameLength > 5 {
		fmt.Println("Your name is too long.")
	} else {
		fmt.Println("Your name is OK.")
	}
}