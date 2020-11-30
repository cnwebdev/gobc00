// The fmt package
package main

import "fmt"

func main() {

	const y = 42

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	const usa = 1776

	fmt.Printf("%#x\n", usa)
}
