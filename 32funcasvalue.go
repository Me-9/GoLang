package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {

	// Function adalah first class citizen(Function bisa disimpan sebagai variable atau dijadikan sebagai value atau nilai)
	// Function juga merupakan tipe data, dan bisa disimpan di dalam variable

	goodbye := getGoodBye //Memanggil func getGoodBye sebagai variable tanpa tuple()
	fmt.Println(goodbye("Man"))

	fmt.Println(getGoodBye("Man")) // Memanggil langsung tanpa menjadikan si func sebagai variable
}
