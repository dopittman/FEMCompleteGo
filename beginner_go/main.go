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

	// Functions

	result := add(3, 4)
	println("this is the result", result)

	sum, product := calculateSumAndProduct(10, 10)

	fmt.Printf("this is sum: %d, and this is product: %d \n", sum, product)

	// Conditionals
	myAge := 30

	// if-else
	if myAge >= 18 {
		fmt.Println("you are an adult")
	} else if myAge >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	// switch - will break once the case is met, no need to do it manually
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("it's the weekend")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	// while loop - does not have a "while" keyword
	counter := 0
	for counter < 3 {
		fmt.Println("This is the counter: ", counter)
		counter++
	}

	// creates an infinite loop - then you can break it when you want with a condition
	iterations := 0
	for {
		if iterations > 3 {
			break
		}
		iterations++
	}

	// Arrays and Slices
	numbers := [5]int{10, 20, 30, 40, 50} // Tells compiler that the array should hold 5 items, also it can not hold different types in same array so all of these must be ints
	fmt.Printf("this is our array: %v\n", numbers)
	fmt.Println("this is the last value: ", numbers[len(numbers)-1])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("this is our matrix:%v\n", matrix)

	// Slices - basically a dynamic array
	// allNumbers := numbers[:] // make a slice array containing all numbers, we can now append more numbers to it
	// firstThree := numbers[0:3] // gets the first three items of numbers

	fruits := []string{"Apple", "Strawberry", "Banana"} // creates a slice by not declaring the Array length, cons: memory allocation, speed, etc
	fmt.Printf("These are my fruits: %v\n", fruits)
	fruits = append(fruits, "Kiwi")
	fmt.Printf("These are my fruits with kiwi: %v\n", fruits)
	fruits = append(fruits, "Mango", "Pineapple")
	fmt.Printf("These are my fruits with more added: %v\n", fruits)

	moreFruits := []string{"Blueberries", "Tomato"}
	fruits = append(fruits, moreFruits...) // can append slices to slices
	fmt.Printf("These are my fruits with EVEN more added: %v\n", fruits)

	// looping through array or slice with ind/val values
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// Maps - map[key type]value type
	capitalCities := map[string]string{
		"USA":   "Washington D.C",
		"India": "New Delhi",
		"UK":    "London",
	}

	fmt.Println(capitalCities["USA"])

	// Check if key exists before accessing it
	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("this is the capital: ", capital)
	} else {
		fmt.Println("Does not exist in map")
	}

	delete(capitalCities, "UK")
	fmt.Printf("This is new deleted map %v\n", capitalCities)

	// Structs

	person := Person{
		Name: "John",
		Age:  25,
	}
	fmt.Printf("this is our person: %v\n", person)
	fmt.Printf("This adds the field names to the struct %+v\n", person)

	// anonymous struct - creates struct in place
	employee := struct {
		name string
		id   int
	}{
		name: "alice",
		id:   123,
	}

	fmt.Println("this is employee: ", employee)

	// nested structs
	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	// You can also omit values, ie. the phone string. It will use the default zero value for the type
	contact := Contact{
		Name: "Mark",
		Address: Address{
			Street: "123 Main Street",
			City:   "Anytown",
		},
	}
	fmt.Println("this is contact: ", contact)

	// This shows how the change is being made to a copy and not the passed data, we must use a pointer(*) in order to pass the reference
	fmt.Println("Name before: ", person.Name)
	modifyPersonName(person)
	fmt.Println("Name after: ", person.Name)

	// need to pass the address of the reference with & attached to the argument
	person.modifyPersonNameWPointer("David")
	fmt.Println("Name after using pointer: ", person.Name)

	x := 20
	ptr := &x // address in memory of pointer
	fmt.Printf("value of x: %d and address of x: %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("value of new x: %d and address of x: %p\n", x, ptr)
}

// -------------------- Functions --------------------

// Making a function capialised makes it exportable/Public, lowercase is not Public
func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

// -------------------- Structs --------------------
type Person struct {
	Name string
	Age  int
}

// Will not modify the actual person being passed in, creates a copy that is deleted after the function finishes
func modifyPersonName(person Person) {
	person.Name = "David"
	fmt.Println("inside scope new name: ", person.Name)

}

// This is a function placed on the Person struct, using a method receiver (p *Person)
func (p *Person) modifyPersonNameWPointer(name string) {
	p.Name = name
	fmt.Println("inside scope new name: ", p.Name)
}
