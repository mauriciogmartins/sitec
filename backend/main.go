package main

import (
	"fmt"
	"sitec/database"
)

func main() {
	db, e := database.Dbconnection()

	if db == nil {
		fmt.Println(e.Error())
	}

	fmt.Println("Conectou")
}
