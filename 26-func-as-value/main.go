package main

import "fmt"

func getGoodbye(name string) string {
	return "Good bye, " + name
}

func main() {
	var thisVarIsFunc func(string) string = getGoodbye
	thisVarIsFuncToo := getGoodbye

	fmt.Println(thisVarIsFunc("Alfian"))
	fmt.Println(thisVarIsFuncToo("Taka"))
}