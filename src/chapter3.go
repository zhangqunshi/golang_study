package main

//import "fmt"  //我才知道不用fmt也能打印

const (
	WIDTH = 100
	HEIGHT = 200
)

func main() {
    const SL, XL, XXL int = 1, 2, 3
	println(SL, XL, XXL);
	println(WIDTH, HEIGHT);

	const (
		a = iota
		b
		c
	)
	println(a, b, c)
}
