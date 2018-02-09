Go语言学习笔记二： 变量
==============
今天又学了一招如何查看go的版本的命令：`go version`。另外上一个笔记中的代码还可以使用`go run hello.go`来运行，只是这种方式不会生成exe文件。

### 定义变量
使用var关键字来定义变量。例如：
```
var x int = 100;
```
顺便说一句，go语言一行代码可以不写分号结束符。
也分单行注释`//`和多行注释`/* */`。
其实也可以不写类型，go语言会推测出数据类型。例如：
```
x ：= 100
```
写法看上去更简单了。还可以写成
```
var x = 200
```
可以一次定义多个变量：
```
var x, y int = 1, 2
```

### 数据类型
类型虽然可以不写，但不代表没有。Go语言具有以下几种数据类型：
1. 布尔型 bool
1. 数字: 整形int和浮点型float32, float64
1. 字符串
1. 派生类型：指针(Pointer), 数组，结构化(struct), Channel类型，函数类型，切片类型，接口类型（interface），Map类型。

Go也有基于架构的类型，例如：int, uint, uintptr。
uint8, uint16, uint32, uint64, int8, int16, int32, int64。

看到go语言支持这么多类型，我都快疯了，这是要干死所有语言的节奏呀（太消耗脑细胞了）。

```
package main

import "fmt"

func main() {
    var a int = 0;
    var b int8 = 127;
    var c int16 = 32767;
    var d int32 = 2147483647;
    var e int64 = 9223372036854775807;
    fmt.Println(a, b, c, d, e);

    var f uint8 = 255;
    var g uint16 = 65535;
    var h uint32 = 4294967295;
    var i uint64 = 18446744073709551615;
    fmt.Println(f, g, h, i);


    var k float32 = 3.1415;
    var l float64 = 3.1415;
    var m complex64 = 3.1415;
    var n complex128 = 3.1415;
    fmt.Println(k, l, m, n);

    w := 1;
    var y bool = true
    z := false
    fmt.Println(w, y, z);
}

```

此系列其他文章地址：
https://github.com/zhangqunshi/golang_study

-------------------------
参考资料：
https://golang.org/doc/
https://go-zh.org/doc/
http://www.runoob.com/go/go-data-types.html

