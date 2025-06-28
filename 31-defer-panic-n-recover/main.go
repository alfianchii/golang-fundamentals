package main

import "fmt"

func logEndApp() {
	fmt.Println("End the app!")
}

func logging() {
	fmt.Println("App is closed!")
}

func runApp(error bool) {
	defer logging()
	defer logEndApp()
	fmt.Println("Run the app...")
	if error {
		panic("ERROR OCCURS!")
	}
}

func main() {
	runApp(true)
}