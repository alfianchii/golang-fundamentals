package main

import (
	"46-package-init/database" // It runs init() first
	_ "46-package-init/internal"
	"fmt"
)

func main() {
	var db string = database.GetDB()
	fmt.Println(db)
}