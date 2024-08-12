package main

import (
	"fmt"
)

func float() {
    // Declare and initialize floating-point numbers
	var a float32 = 3.14                          // 32-bit floating-point
	var b float64 = 2.718281828459045             // 64-bit floating-point

	// Perform some basic operations with floating-point numbers
	sum := a + float32(b)                         // Sum of a and b, converting b to float32
	product := a * float32(b)                     // Product of a and b, converting b to float32
	difference := b - float64(a)                  // Difference, converting a to float64
	quotient := b / float64(a)                    // Quotient, converting a to float64

	// Print results for floating-point numbers
	fmt.Println("Float32 a:", a)
	fmt.Println("Float64 b:", b)
	fmt.Println("Sum of a and b:", sum)
	fmt.Println("Product of a and b:", product)
	fmt.Println("Difference (b - a):", difference)
	fmt.Println("Quotient (b / a):", quotient)

	// Declare and initialize complex numbers
	var c complex64 = 1 + 2i                       // 32-bit complex number
	var d complex128 = 3 + 4i                      // 64-bit complex number

	// Perform basic operations with complex numbers
	complexSum := c + complex64(d)                // Sum of c and d, converting d to complex64
	complexProduct := c * complex64(d)            // Product of c and d, converting d to complex64

	// Combining floating-point and complex numbers
	combinedSum := complex(float32(a), 0) + c     // Combine float32 a with complex number c
	combinedProduct := complex(float64(b), 0) * d  // Combine float64 b with complex number d

	// Print results for complex numbers
	fmt.Println("Complex64 c:", c)
	fmt.Println("Complex128 d:", d)
	fmt.Println("Sum of c and d:", complexSum)
	fmt.Println("Product of c and d:", complexProduct)

	// Print results for combined operations
	fmt.Println("Combined Sum (a + c):", combinedSum)
	fmt.Println("Combined Product (b * d):", combinedProduct)
}
