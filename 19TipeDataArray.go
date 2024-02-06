package main

import "fmt"

func main()  {

	//Contoh array string
	var names[3] string
	names[0] = "Emma"
	names[1] = "Falls"
	names[2] = "In"

	fmt.Println(names)
	// fmt.Println(names[1])
	// fmt.Println(names[2])


	//Contoh array int
	var values = [3]int{
		90,
		91,
		92,
	}
	
	fmt.Println(values)


	//Array tanpa menentukan panjangnya 
	var val = [...]int{
		10,
		20,
		30,
		40,
	}

	fmt.Println(val)
	fmt.Println(len(val))
}