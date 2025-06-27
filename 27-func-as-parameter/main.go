package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	var filteredName string = filter(name)
	fmt.Println("Hello, " + filteredName + "!")
}

func spamFilter(word string) string {
	if word == "Fuck" {
		return "****"
	} else {
		return word
	}
}

type Filter func(string) string

func sayHelloWithFilterPart2(name string, filter Filter) {
	var filteredName string = filter(name)
	fmt.Println("Hello, " + filteredName + "!")
}

func main() {
	var myName string = "Alfian Taka"
	sayHelloWithFilter(myName, spamFilter)

	var comment string = "Fuck"
	var filterWord func(string) string = spamFilter
	sayHelloWithFilter(comment, filterWord)

	var myRealName string = "Alfian Handsome"
	var filterName Filter = filterWord
	sayHelloWithFilterPart2(myRealName, filterName)
}