// Interface
// Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung 
// Sebuah interface berisikan definisi-definisi method
// Biasanya interface digunakan sebagai kontrak atau prototype

// Implementasi Interface
// Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
// Sehingga kita tidak perlu mengimplementasikan interface secara manual
// Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main()  {
	man := Person{Name: "Man"}
	SayHello(man)

	kucing := Animal{Name: "Kucing"}
	SayHello(kucing)
}