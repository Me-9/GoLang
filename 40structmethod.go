// Struct Method
// Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
// Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
// Method adalah function
// Lanjut juga di part 39

package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, my name is", customer.Name)
}

func main() {
	man := Customer{Name: "Man"}
	man.sayHello()
}
