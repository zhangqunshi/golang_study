Go语言学习笔记八： 数组
==============
数组地球人都知道。所以只说说Go语言的特殊（奇葩）写法。
我一直在想一个人参与了两种语言的设计，但是最后两种语言的语法差异这么大。这是自己否定自己么，为什么不与之前统一一下。

### 声明数组
```
var variable_name [SIZE] variable_type
```
例子：
```
var x [10] int
```
### 初始化数组
```
var x = [5] int {1, 2, 3, 4, 5}
var y = [...] int {1, 2, 3, 4, 5}
```
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小.

### 多维数组
```
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

var x [5][10][4]int
```
初始化多维数组
```
a = [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11}   /*  第三行索引为 2 */
}
```

### 向函数传递数组
```
void myFunction(param [10]int) {
}
或者
void myFunction(param []int) {
}
```

此系列其他文章地址：
https://github.com/zhangqunshi/golang_study
