// Type Assertions
// Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering kali digunakan ketika kita bertemu dengan tipe data interface kosong

package main

import "fmt"

func random() interface{} {
	return true
}


// Type Assertions Menggunakan Switch
// Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
// Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
// Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions

func main()  {
	result := random()

	// Silakan komen 2 baris kode di bawah jika return dalam func random() tipe datanya adalah selain string
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // Akan terjadi panic
	// fmt.Println(resultInt)

	// Type Assertions Switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}