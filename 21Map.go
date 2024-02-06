package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Buat map baru dan delete salah satu key nya
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Man"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
