package main

import (
	"fmt"
	"math"
)

func main() {
	const fact = "mac and cheese is good"
	// fact = "no it isnt" won't work because const vars cannot be reassigned
	var a = math.Sqrt(4) // works
	fmt.Println("a:", a)
	// const b = math.Sqrt(4) // wont work because value won't be known at compile time
	const str = "nya"
	// str is untyped, but will provide its default type if the code explicitly requests it
	fmt.Printf("str: type %T | value %v", str, str)
	// To create a typed constant:
	const typedfact2 string = "I have a crush :flushed:"
}