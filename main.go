package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Emily's cuteness measuring program! Nyaa!")
	fmt.Println()
	fmt.Println("Enter name:")
	text, _ := reader.ReadString('\n')
	fmt.Println(text, "is 2000% cute!")
}
