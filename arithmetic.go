package main

import "fmt"

func Arithmetic() {
	a := 10; b := 5; tambah := a + b; kurang := a - b
	kali := a * b; bagi := a / b; modulus := a % b

	fmt.Printf("%d + %d = %d\n", a, b, tambah)
	fmt.Printf("%d - %d = %d\n", a, b, kurang)
	fmt.Printf("%d * %d = %d\n", a, b, kali)
	fmt.Printf("%d / %d = %d\n", a, b, bagi)
	fmt.Printf("%d %% %d = %d\n", a, b, modulus)

	a++; fmt.Printf("Increment: %d\n", a) // Increment
	b--; fmt.Printf("Decrement: %d\n", b) // Decrement
}