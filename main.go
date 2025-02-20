package main

import (
	"fmt"
)

func main() {
	var bit uint8 = 0b1111_1111
	fmt.Println("Hello, World!")
	fmt.Printf("overflowed bit is %08b\n", bit+1)
}
