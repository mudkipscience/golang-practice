package main

import "fmt"

func main() {
	// Booleans (bool)
	a := true
	b := false

	fmt.Println("a:", a, "b:", b)

	c := a && b // Returns true when both variables are true
	fmt.Println("c:", c)

	d := a || b // Returns true if either variable is true
	fmt.Println("d:", d)
	fmt.Println()
	fmt.Println("%a")
}

/* https://golangbot.com/types/

Unsigned int: Holds large positive value, but no negative values (i.e 0 to 255)
Signed int: Hold both positive and negative values (i.e 0 to +127, -1 to +128)

int will represent a 32bit or 64bit integer depending on the platform

range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

Types
	bool
	Numeric Types
		int8, int16, int32, int64, int
        uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
		byte
		rune
	string
*/
