// Loop example 01
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i <= 9 {
			fmt.Printf("0%d :", i)
		} else {
			fmt.Printf("%d :", i)
		}
		fmt.Printf("Hello, GoLang Loop!\n")
	}
}
