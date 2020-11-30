// Zero value
package main

import "fmt"

func main() {
	var a string
	var b int
	var c bool
	var d int64
	var e uint
	var f float64

	fmt.Printf("Zero value of a = %s type %T\n", a, a)
	fmt.Printf("Zero value of b = %d type %T\n", b, b)
	fmt.Printf("Zero value of c = %v type %T\n", c, c)
	fmt.Printf("Zero value of d = %d type %T\n", d, d)
	fmt.Printf("Zero value of e = %d type %T\n", e, e)
	fmt.Printf("Zero value of f = %f type %T\n", f, f)
}
