package main

import "fmt"

func main() {

	//contoh 1 = mendeklarasikan tipe data suatu variable
	var name string

	name = "Man"
	fmt.Println(name)

	name = "Nam"
	fmt.Println(name)

	// Contoh 2 = Langsung tanpa harus sebut tipe data (exam = string, int, etc)
	var nama = "KL"
	fmt.Println(nama)

	nama = "JKT"
	fmt.Println(nama)

	// Contoh 3 = Lebih simpel lagi (menggunakan := untuk mendeklarasikan suatu variable)
	dia := "Komang"
	fmt.Println(dia)

	dia = "Kaming"
	fmt.Println(dia)

	// Contoh 4 = Deklarasi multiple variable
	var (
		programName  = "GoProj"
		codeLanguage = "Go"
		Level        = "Beginner"
	)

	fmt.Println(programName, codeLanguage, Level)
}
 