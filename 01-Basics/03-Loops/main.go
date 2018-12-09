package main

import "fmt"

func main() {
	// For Loop
	for i:=0; i<100; i++ {
	fmt.Printf("%d - %b - %x - %o - %#x - %q \n", i, i, i, i, i, i)	 //Prints 0-99 in Decimal, Binary, Hex, Octal, UTF-8
	}
}
