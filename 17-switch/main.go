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
}