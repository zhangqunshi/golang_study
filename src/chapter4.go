package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	var c int
	c = a + b
	fmt.Printf("c=%d\n", c)

	var ptr *int
	ptr = &a
	fmt.Printf("prt=%d\n", *ptr)
}