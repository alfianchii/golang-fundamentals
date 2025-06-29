package main

import (
	"46-package-init/database" // It runs init() first
	"fmt"
)

func main() {
	var db string = database.GetDB()
	fmt.Println(db)
}