package main

import "fmt"

func main() {
	var name string = "Taka"

	switch name {
	case "Alfian":
		fmt.Println("Halo, " + name)
	case "Taka":
		fmt.Println("こんにちは, " + name)
	case "Ogi":
		fmt.Println("Hello, " + name)
	case "Taufan":
		fmt.Println("สวัสดี, " + name)
	default:
		fmt.Println("Hello! What's ur name?")
	}

	// Short
	switch nameLength := len(name); nameLength > 5 {
	case true:
		fmt.Println("Your name is too long.")
	case false:
		fmt.Println("Your name is OK.")
	}

	// Without condition (better use if)
	length := len(name);
	switch {
	case length > 10:
		fmt.Println("Your name is very long.")
	case length > 5:
		fmt.Println("Your name is too long.")
	default: 
		fmt.Println("Your name is OK.")
	}
}