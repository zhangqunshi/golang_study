package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
	a, b := swap("hello", "kris")
	fmt.Println(a, b)

	x := 2
	y := 3
	swap2(&x, &y)
	fmt.Println(x, y)
	
	next := getSequence()
    fmt.Println(next())
    fmt.Println(next())
}
