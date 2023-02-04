package main

import (
	"fmt"

	"rsc.io/quote"
)

// variables outside the function
var myVariable13 = 1
var myVariable14 string = "Sukant"
var myVariable15 bool

// below is not possible
// myVariable16 := 2

func main() {
	HelloWorld()
	ConsumingImportedModule()
	WorkingWithVariables()
}

// simple hello world function
func HelloWorld() {
	fmt.Println("Hello World")
}

// to understand how to consume functions in another package
func ConsumingImportedModule() {
	fmt.Println(quote.Go())
}

// to understand how to work with variables
func WorkingWithVariables() {

	// Declaraing and initializing variables explicitly
	var myVariable1 int = 1
	var myVariable2 string = "Sukant"
	var myVariable3 bool = true

	// Declaring and initializing variables at Go's mercy (implicitly)
	var myVariable4 = 1
	var myVariable5 = "Sukant"
	var myVariable6 = true

	// declaring and initializing variables implicitly, go shortcut
	myVariable7 := 1
	myVariable8 := "Sukant"
	myVariable9 := true

	// declaring but not initializing variables
	var myVariable10 int
	var myVariable11 string
	var myVariable12 bool

	// printing the explicitly declared and initialized variables
	fmt.Println(myVariable1)
	fmt.Println(myVariable2)
	fmt.Println(myVariable3)
	fmt.Println(myVariable4)
	fmt.Println(myVariable5)
	fmt.Println(myVariable6)
	fmt.Println(myVariable7)
	fmt.Println(myVariable8)
	fmt.Println(myVariable9)
	fmt.Println(myVariable10)
	fmt.Println(myVariable11)
	fmt.Println(myVariable12)
	fmt.Println(myVariable13)
	fmt.Println(myVariable14)
	fmt.Println(myVariable15)
}


