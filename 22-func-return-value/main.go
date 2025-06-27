package main

import "fmt"

func getGreetings(name string) string {
	var greetings string = "Hello, " + name + "!"
	return greetings
}

func main() {
	var name string = "Alfian Taka"
	var greetingsToName string = getGreetings(name)
	
	fmt.Println(greetingsToName)
	fmt.Println(getGreetings("Moe"))
	fmt.Println(getGreetings("Galih"))
}