package main

import (
	"fmt"
)

func integer() {
    // Declare and initialize integers of different types
    var a int = 42          // Default int (usually 64 bits on 64-bit systems)
    var b int8 = 12         // 8-bit signed integer
    var c int16 = -1000     // 16-bit signed integer
    var d int32 = 100000    // 32-bit signed integer
    var e int64 = 1234567890123456789 // 64-bit signed integer

    // Perform some basic operations
    sum := a + int(b) + int(c) + int(d) + int(e) // Sum of all integers
    // difference := e - d                       // will cause mismatch error
	difference := e - int64(d)                   // Convert d to int64 for subtraction

    // Print the results
    fmt.Println("Integer a:", a)
    fmt.Println("Integer b:", b)
    fmt.Println("Integer c:", c)
    fmt.Println("Integer d:", d)
    fmt.Println("Integer e:", e)
    fmt.Println("Sum of all integers:", sum)
    fmt.Println("Difference between e and d:", difference)
}
