package main

import "fmt"

func main() {
	var x map[string]string
	x = make(map[string]string)
	
	x["a"] = "1";
	x["b"] = "2";
	
	for item := range x {
		fmt.Println(item, x[item]);
	}
	
	value, exist := x["a"]
	if (exist) {
		fmt.Println("x has", value);
	}
}