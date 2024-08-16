package main

import (
	"fmt"
	"testing"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	// without testify
	if result != 40 {
		// the first way
		t.Fail()
		// the second way
		// t.FailNow()
		// the third way
		// t.Error("Result has to be 40")
	}

	// with testify
	// require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	// without testify
	if result != 20 {
		// panic("Result should be 20")
		t.FailNow()
	}

	// with testify
	// assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}
