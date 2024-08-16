package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func FunctionWithReturnValue() {
	result := getHello("Alice") // Output: Hello Alice
	fmt.Println(result)
}
