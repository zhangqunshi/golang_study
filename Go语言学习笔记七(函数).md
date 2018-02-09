Go语言学习笔记七： 函数
==============

Go语言有函数还有方法，神奇不。这有点像python了。

### 函数定义
```
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
举个例子：
```
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}
```

### 函数调用
```
var a int = 1
var b int = 2
var c int
c = max(a, b)
```

### 函数可以返回多个值
```
func swap(x, y string) (string, string) {
   return y, x
}

func main() {
    a, b := swap("hello", "kris")
    fmt.Println(a, b)
}
```
是不是有点像python.

### 值传递
传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

基础类型都是值传递，地球人都知道。

### 引用传递
引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
```
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保持 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```

天呀，我感觉自己又回到了上学时代。（感觉自己好年轻）

### 函数作为值
Go 语言可以很灵活的创建函数，并作为值使用。
```
func main(){
   /* 声明函数变量 */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* 使用函数 */
   fmt.Println(getSquareRoot(9))

}
```

是不是感觉像javascript语言。

### 闭包
Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
```
func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    next := getSequence()
    fmt.Println(next())
    fmt.Println(next())
}
```

这个写法太。。。完美（真赶时髦）。我猜 `func() int`是闭包的返回值。

### 方法
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
```
func (variable_name variable_data_type) function_name() [return_type]{
   /* 函数体*/
}
```
举个例子:
```
/* 定义结构体 */
type Circle struct {
  radius float64
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```
感觉写法很别扭，Circle给我的感觉就是一个Java的类，而getArea()是这个类的方法。估计只要在函数前面加上`(c Circle)`的都是属于c的方法。




此系列其他文章地址
----------
https://github.com/zhangqunshi/golang_study

参考： http://www.runoob.com/go/go-method.html

