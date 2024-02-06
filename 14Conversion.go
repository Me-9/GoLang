package main

import "fmt"

func main() {
	var nilai32 int32 = 37824
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Man"
	var e = name[0]
	var eString = string(e)
	var ei uint8 = 16
	var eStringEi = string(ei)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
	fmt.Println(ei)
	fmt.Println(eStringEi)
	
}