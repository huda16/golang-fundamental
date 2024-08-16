package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func FunctionWithParameter() {
	sayHelloTo("Alice", "Bob") // Output: Hello, Alice Bob
}
