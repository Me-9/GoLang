package main

import (
	"fmt"

)

// Recover
// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

// Recover yang salah
func endApp()  {
	fmt.Println("End App")
}

func runApp(error bool)  {
	defer endApp()

	if error {
		panic("ERROR")
	}

	message := recover()
	fmt.Println("Terjadi error", message)
}


// Recover yang benar
// Recover disimpan di function yang kita panggil menggunakan defer (contoh "defer endApp2")

func endApp2()  {
	fmt.Println("END APP")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp2(error bool)  {
	defer endApp2()
	if error {
		panic("Error")
	}
}

func main()  {
	runApp(false)
	runApp2(true)
	fmt.Println("Man")
}