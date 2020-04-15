package main

import (
	"fmt"
	"unsafe" // Used for the Sizeof() function. Can cause portability issues, use with care!
)

func main() {
	// Booleans (bool)
	a, b := true, false
	
	fmt.Println("Booleans:")
	fmt.Println("a:", a, "b:", b)

	c := a && b // Returns true when both variables are true
	fmt.Println("c:", c)

	d := a || b // Returns true if either variable is true
	fmt.Println("d:", d)
	fmt.Println()

	e, f := 69, 420

	fmt.Println("Signed integers:")
	// Variable type can be printed with %T, %d to print size
	fmt.Printf("Type of e: %T | Size of e: %d", e, unsafe.Sizeof(e)) // Type should be int, size should be 8 (or 4 if ran on 32bit system)
	fmt.Printf("\nType of f: %T | Size of f: %d", f, unsafe.Sizeof(f))
	fmt.Println()

	fmt.Println("Floating point types:")
	g, h := 13.37, 4.20
	fmt.Printf("Type of g: %T | Type of h: %T\n", g, h)
	sum := g + h
	diff := g - h
	fmt.Println("Sum of g + h:", sum, "| Difference:", diff)
	fmt.Println()

	// i dont really know what im doing here complex numbers are kinda confusing
	// maybe i should ask my maths teacher
	c1 := complex(5, 7) // Complex (5 real, 7 imaginary)
	c2 := 8 + 27i // Complex (8 real, 27 imaginary)
	cadd := c1 + c2
	cmul := c1 * c2
	fmt.Println("Complex numbers:")
	fmt.Println("Value of c1:", c1, "| Value of c2:", c2)
	fmt.Println("sum:", cadd)
	fmt.Println("product", cmul)
	fmt.Println()

	str := "sans"
	str2 := "undertale"
	joined := str + " " + str2 // join str and str2, duh
	fmt.Println("Strings:")
	fmt.Println("str:", str)
	fmt.Println("str2:", str2)
	fmt.Println("Joined:", joined)
	fmt.Println()

	fmt.Println("Type conversion:")
	/* This won't work, because go is strict about explicit typing
	    i := 55      //int
    	j := 67.8    //float64
    	sum := i + j //int + float64 not allowed
		fmt.Println(sum)
	- j must be converted to an integer before we can add
	*/
	i := 55
	j := 67.8
	sum2 := i + int(j) // converts j to int
	fmt.Println("Value of i:", i, "| Value of j:", j)
	fmt.Println("Value of j converted to int:", int(j))
	fmt.Println("sum (after converting j to int):", sum2)
	k := 10
	fmt.Printf("Value of k: %v | Type of k: %T", k, k)
	l := float64(k) // Convert k to a float
	fmt.Printf("\nk converted to float64: Value %v | Type %T\n", l, l)
	fmt.Println("Sum of j + k (l)", j + l)
}
