package main

import (
	"fmt"
)

func main() {
	var name string // Declares a variable named age of type string
	name = "Emily"  // Assigns a value

	var age = 16 // You don't need to specify a type if you assign a value when the variable is declared

	height := "5'10\"" // Shorthand declaration

	var coolnumber, coolnumber1, coolnumber2 int = 420, 69, 1337 // Declaring two vars at once woo (type here can be omitted)

	var (
		favouriteThing = "mac and cheese"
		cutenessLevel  int // Equals 0
	) // Declaring multiple vars of different types

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("My height is", height)
	fmt.Println("Cool numbers:", coolnumber, coolnumber1, coolnumber2)
	fmt.Println("My favourite thing is", favouriteThing)
	fmt.Println("My cuteness level is", cutenessLevel, "%")
}

/*
Notes:
- Variables of one type can't be declared as another (var int cannot be assigned a string value)
- Shorthand declaration only works when one of the variables is newly declared
*/
