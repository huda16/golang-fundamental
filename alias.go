package main

import (
	"fmt"
)

func alias() {
    // Declare variables of different basic data types
    var a byte = 255                  // byte (uint8)
    var b rune = 'G'                  // rune (int32), representing the Unicode code point for 'G'
    var c int = -42                   // int (signed integer)
    var d uint = 42                   // uint (unsigned integer)

    // Display the values and types of each variable
    fmt.Printf("Value of byte a: %d, Type: %T\n", a, a)
    fmt.Printf("Value of rune b: %c, Type: %T\n", b, b) // %c to print the character
    fmt.Printf("Value of int c: %d, Type: %T\n", c, c)
    fmt.Printf("Value of uint d: %d, Type: %T\n", d, d)

    // Perform some operations
    var sum = int(a) + c          // Sum of byte and int (convert byte to int)
    var difference = int(d) - c   // Difference between uint and int (convert uint to int)

    // Print results of operations
    fmt.Printf("Sum of a and c: %d\n", sum)
    fmt.Printf("Difference between d and c: %d\n", difference)

    // Example of converting rune to string
    var str string = string(b)     // Convert rune to string
    fmt.Printf("Rune b as string: %s\n", str)
}
