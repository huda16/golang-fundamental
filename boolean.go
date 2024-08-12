package main

import (
	"fmt"
)

func Boolean() {
	// Declare and initialize boolean variables
	var isRaining bool = false
	var isWeekend bool = true

	// Display initial values
	fmt.Println("Is it raining?", isRaining)
	fmt.Println("Is it the weekend?", isWeekend)

	// Perform logical operations
	canGoOutside := !isRaining && isWeekend // Can go outside if it's not raining and it's the weekend
	fmt.Println("Can I go outside?", canGoOutside)

	// Change the state of isRaining
	isRaining = true
	fmt.Println("Is it raining now?", isRaining)

	// Re-evaluate the ability to go outside
	canGoOutside = !isRaining && isWeekend
	fmt.Println("Can I go outside now?", canGoOutside)

	// Use boolean in a conditional statement
	if canGoOutside {
		fmt.Println("Great! Let's go out!")
	} else {
		fmt.Println("Maybe we should stay indoors.")
	}

	// Example of using boolean expressions
	age := 20
	hasPermission := true

	// Check if the person can enter a restricted area
	canEnter := age >= 18 && hasPermission
	fmt.Println("Can the person enter the restricted area?", canEnter)
}
