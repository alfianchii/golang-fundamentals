package helper

import "fmt"

var version = "1.0.0"
var App = "golang"

func sayGoodBye(name string) string {
	return "Goodbye, " + name + " :("
}

func SayHello(name string) string {
	return "Hello, " + name + "!"
}

func Example() {
	var result string = sayGoodBye("Galih")
	fmt.Println(version, result)
}