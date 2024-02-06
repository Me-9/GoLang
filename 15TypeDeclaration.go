package main

import "fmt"

func main()  {
	type noKTP string

	var ktpMan noKTP = "11111"

	var contoh string = "22222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpMan)
	fmt.Println(contohKTP)

	for o := 0; o < 127; o++ {
		println(string(o))
	}
}