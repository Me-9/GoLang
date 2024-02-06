package main

import "fmt"

func main()  {
	var nilaiAkhir = 90
	var nilaiAbsensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusNilaiAbsensi bool = nilaiAbsensi > 70

	var lulus bool = lulusNilaiAbsensi && lulusNilaiAkhir

	fmt.Println(lulus)
}