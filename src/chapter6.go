package main

import "fmt"

func main() {

	for a := 0; a < 10; a++ {
		fmt.Printf("a=%d\n", a)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
