package main

import "fmt"

func main() {
	nums := []int {10, 20, 30}
	sum := 0
	
	for i, num := range nums {
		fmt.Println(i, num);
		sum += num
	}
	
	fmt.Println("sum: ", sum);
}