// Exploring Type
package main

import "fmt"

func main() {

	var x = 52
	var y string

	y = "Thanh Pho Saigon"
	// y = x // y is type string, x is type int, therefore x can not be assigned to y.

	fmt.Printf("%d %T\n", x, x)
	fmt.Printf("%s %T\n", y, y)

}
