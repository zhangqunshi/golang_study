Go语言学习笔记十： 结构体
==============
Go语言的结构体语法和C语言类似。而结构体这个概念就类似高级语言Java中的类。

### 结构体定义
结构体有两个关键字type和struct，中间夹着一个结构体名称。大括号里面写上所有的成员变量，并且指定这些变量的类型。访问这些内部成员时使用.符号。注意是结构体创建的变量，才能用点访问内部成员。
不是直接用结构体访问。

```
package main

import "fmt"

type Book struct {
	name string
	price int
}

func main() {
	var book1 Book;
	var book2 Book;
	
	book1.name = "书名1"
	book1.price = 100
	
	book2 = Book {"书名2", 200}
	
	fmt.Printf( "Book 1 name : %s\n", book1.name)
	fmt.Printf( "Book 1 price : %d\n", book1.price)
	fmt.Printf( "Book 2 : %s\n", book2)
}
```


### 结构体作为函数参数

```
func printBook( book Book ) {
   fmt.Printf( "Book name : %s\n", book.name);
   fmt.Printf( "Book price : %d\n", book.price)
}
```

### 结构体指针
```
var x *Book
x = &book1;
x.name;
```

结构体指针和结构体变量都是使用.来访问内部成员。感觉使用指针操作更麻烦一些。




