package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{
	// 	"name": "Alfian",
	// 	"address": "Indonesia",
	// }
	// var person = map[string]string{
	// 	"name": "Alfian",
	// 	"address": "Indonesia",
	// }
	person := map[string]string{
		"name": "Alfian",
		"address": "Indonesia",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Funcs
	var book map[string]string = make(map[string]string)
	book["title"] = "Journey of My Life"
	book["author"] = "Alfian Taka"
	book["apkhBtul"] = "ia"

	delete(book, "apkhBtul")

	fmt.Println(book)
	fmt.Println(len(book))
}