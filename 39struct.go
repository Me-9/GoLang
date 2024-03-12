// Struct
// Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// Struct biasanya representasi data dalam program aplikasi yang kita buat
// Data di struct disimpan dalam field
// Sederhananya struct adalah kumpulan dari field

// Membuat Data Struct
// Struct adalah template data atau prototype data
// Struct tidak bisa langsung digunakan
// Namun kita bisa membuat data/object dari struct yang telah kita buat

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct method
func (customer Customer) sayHello(name string)  {
	fmt.Println("Hello", name,", my name is", customer.Name)
}

func main() {
	var wisman Customer

	// Cara pertama
	wisman.Name = "Wisman"
	wisman.Address = "Indonesia"
	wisman.Age = 18

	// Cara kedua
	fmt.Println(wisman)
	fmt.Println(wisman.Name)
	fmt.Println(wisman.Address)
	fmt.Println(wisman.Age)

	fmt.Println(wisman)

	man := Customer{"Wisman", "Indonesia", 18}
	fmt.Println(man)

	wisman.sayHello("All of the people") // Output struct method
	
}
