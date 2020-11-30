// Short declaration operator
package main

import "fmt"

func main() {
	// Short declaration of a, b, c and p variables by using := operator
	// Short declaration should only be done inside a code block,
	// not recommended for package level declaration
	a := "The United States "
	b := "of America"

	p := 331002651

	c := a + b
	fmt.Println(c)
	fmt.Printf("The United States population is %d\n", p)
}
