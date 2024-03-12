package main

import "fmt"

// Closure adalah kemampuan sebuah function berinteraksi dengan data-data di sekitarnya dalam scope yang sama
// Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

func main()  {
	counter := 0

	increment := func ()  {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}