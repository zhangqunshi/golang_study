Go语言学习笔记十三： Map集合
==============

Map在每种语言中基本都有，Java中是属于集合类Map，其包括HashMap, TreeMap等。而Python语言直接就属于一种类型，写法上比Java还简单。
Go语言中Map的写法比Java简单些，比Python繁琐。

### 定义Map
```
var x map[string]string
x : = make(map[string]string)
```
写法上有些奇怪，map为关键字，右侧中括号内部为key的类型，中括号外部为value的类型。一般情况下使用逗号或者冒号分割key和value，但是Go语言没有遵循这个原则，而是使用了括号里和括号外。
而且map必须初始化，否则就会变成nil map，而nil map不能用来存放键值对。

```
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
```

### delete函数
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。

```
fruits := map[string]string {"apple": "12", "orange": "210"}
delete(fruits, "apple")
```
