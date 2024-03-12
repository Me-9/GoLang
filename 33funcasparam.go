package main

import "fmt"

// Function tidak hanya bisa kita simpan di dalam variable sebagai value
// Namun juga bisa kita gunakan sebagai parameter untuk function lain


// Function Type Declarations
// Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

func sayHelloWithFilter(name string, filter Filter) { //Filter nya dipanggil di parameter filter di samping deh (ga paham konteksnya? ke bawah dulu)
	filteredName := filter(name)

	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	// sayHelloWithFilter("Anjing", )
}

// Function Type Declarations
type Filter func(string) string //Simplify
// Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter