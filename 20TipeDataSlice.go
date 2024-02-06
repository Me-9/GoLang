package main

import "fmt"

func main()  {
	names := [...]string{"M", "a", "n"}

	slice1 := names[:]
	fmt.Println(slice1)

	slice := names[1:2]
	fmt.Println(slice)

	slice2 := names[2:3]
	fmt.Println(slice2)


	// Append Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)


	// New Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eki")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)


	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}