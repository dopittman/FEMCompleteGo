package main

import (
	"fmt"
)

func main() {

	// -------------------- Variables and Types --------------------
	var name string = "David"
	fmt.Printf("Hello, %s\n", name)

	age := 35                               // Short variable declaration with type inference
	fmt.Printf("this is my age: %d\n", age) // Using %d for integers

	var city string
	city = "Charlotte"
	fmt.Printf("this is my city %s\n", city)

	var country, continent string = "USA", "North America"

	fmt.Printf("this is my country %s and this is my continent %s\n", country, continent)

	// Multi variable declaration and assignment
	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)

	fmt.Printf("isEmployed: %t, this is my salary: %d, and this is my position: %s\n", isEmployed, salary, position)

	// Zero values - will avoid being null, but can lead to unwanted conditions when used in code
	var defaultInt int       // 0
	var defaultFloat float64 // 0.000000
	var defaultString string // ""
	var defaultBool bool     // false

	fmt.Printf("int: %d, float: %f, string: %s, and bool: %t", defaultInt, defaultFloat, defaultString, defaultBool)

	//constants - don't change
	const pi = 3.14 // Compiler doesn't care about unused const due to memory allocation. If string is unused, it doens't want to allocate memory to it. const gets allocated regardless
	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
	)
	fmt.Printf("Monday: %d, Tuesday: %d, Wednesday: %d\n", monday, tuesday, wednesday)

	const typedAge int = 25
	const untypedAge = 25 // untyed const

	fmt.Println(typedAge == untypedAge) // returns true

	// Enums-ish using iota
	const (
		Jan = iota + 1 // 1
		Feb            // 2
		Mar            // 3
	)

	fmt.Printf("jan: %d, feb: %d, mar: %d\n", Jan, Feb, Mar)

	result := add(3, 4)
	println("this is the result", result)

	sum, product := calculateSumAndProduct(10, 10)

	fmt.Printf("this is sum: %d, and this is product: %d \n", sum, product)
}

// -------------------- Functions --------------------

// Making a function capialised makes it exportable/Public, lowercase is not Public
func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
