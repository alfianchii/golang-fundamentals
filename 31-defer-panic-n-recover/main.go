package main

import "fmt"

func logEndApp() {
	fmt.Println("End the app...")
	msg := recover()
	fmt.Println("Error occurs:", msg)
}

func logging() {
	fmt.Println("App is closed.")
}

func runApp(error bool) {
	// Deferred funcs run in LIFO Order
	defer logging()
	defer logEndApp()
	fmt.Println("Run the app...")
	if error {
		panic("WOOPS!")
	}
}

func main() {
	runApp(true)
	fmt.Println("We passed the panic disaster!")
}