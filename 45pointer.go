// Pass by Value
// Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
// Artinya, jika kita mengirim sebuah variable ke dalah function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value-nya

package main

import "fmt"

type Address struct {
	City, Province, Country string
}


// Pointer
// Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menDUPLIKASI data yang sudah ada
// Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

// Operator &
// Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable-nya


func main()  {
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}

	address2 := address1 // menDUPLIKAT address1, bukan menggunakan (Masing-masing mempunyai value-nya tersendiri)
	address2.City = "Bandung"

	fmt.Println(address1) // address1 tetap tidak akan berubah output-nya
	fmt.Println(address2) // Berubah menjadi Bandung untuk CITY-nya

	// Kesimpulannya
	// Apa yang diubah oleh address2, tidak akan merubah si address1


	// Operator &
	address3 := &address1 // Kita menggunakan (&) untuk mengarah ke address1 tanpa menDUPLIKASInya
	address3.City = "Sukabumi"

	fmt.Println(address1) // Merubah value city di address1 menjadi Sukabumi
	fmt.Println(address3)
}