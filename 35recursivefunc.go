package main

import "fmt"

// Recursive function adalah function yang memanggil function dirinya sendiri
// Kadang dalam pekerjaan, kita sering menemui kasus di mana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
// Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial dalam Matematika

// Contoh tidak menggunakan Function Recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	
	return result
}

// Contoh menggunakan Function Recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))       // Tidak menggunakan Func Recursive
	fmt.Println(factorialRecursive(10)) // Menggunakan Func Recursive
}
