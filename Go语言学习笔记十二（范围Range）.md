Go语言学习笔记十二： 范围(Range)
==============

rang这个关键字主要用来遍历数组，切片，通道或Map。在数组和切片中返回索引值，在Map中返回key。
这个特别像python的方式。不过写法上比较怪异使用`:=`分割，而在python中使用`in`分割。而python中range是函数，不是关键字。

```
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
```