package main

import "fmt"

func main() {
	var x = [5] int {1, 2, 3, 4, 5}
	var y = [] int {1, 2, 3, 4, 5}
	var z = [...] int {1, 2, 3, 4, 5}
	fmt.Println(x, y, z)
}
