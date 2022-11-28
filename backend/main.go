package main

import (
	"fmt"
	"sitec/enuns"
)

func main() {
	var c = enuns.Instagram

	fmt.Println(c)
	fmt.Println(c.ContactsTypeToString())
}
