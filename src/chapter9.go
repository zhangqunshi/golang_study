package main

import "fmt"

func main() {
	var ptr1 *int
	var ptr2 *float32

	var a int = 10

	ptr1 = &a;

	fmt.Printf("a变量的地址: %x\n", &a); 

	fmt.Printf("ptr1变量保存的地址: %x\n", ptr1);

	fmt.Printf("ptr1变量地址指向的值: %d\n", *ptr1)
	
	fmt.Printf("ptr2变量的值: %x\n", ptr2)
}