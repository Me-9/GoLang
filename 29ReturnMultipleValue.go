package main

import "fmt"

func getFullName() (string, string) {
	return "Roy", "Wisman"
}

func main()  {

	// Return semuanya
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)


	// Return hanya yang diinginkan
	firstName, _ = getFullName()
	fmt.Println(firstName)

	_, lastName = getFullName()
	fmt.Println(lastName)
}