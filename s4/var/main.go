// The var keyword
package main

import "fmt"

func main() {

	// var keyword declaration
	var (
		fname string
		lname string
		age   int
	)

	fname = "Chris"
	lname = "Nguyen"
	age = 50

	fmt.Printf("%s %s is %d year old\n", fname, lname, age)
}
