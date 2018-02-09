Go语言学习笔记五： 条件语句
==============

### if语句
```
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}
```
竟然没有括号，和python很像。但是有大括号，与python又不一样。

例子：
```
package main

import "fmt"

func main() {
   var a int = 1
   if a < 2 {
       fmt.Printf("a < 2\n" )
   }
   fmt.Printf("a = %d\n", a)
}
```

### if..else语句
```
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}
```

### if..else if..else 语句嵌套
```
if score >= 90 {
    fmt.Println("优秀")
} else if score >= 70 {
    fmt.Println("良好")
} else if score >= 60 {
    fmt.Println("一般")
} else {
    fmt.Println("差")
}
```

### if 语句嵌套
```
if 布尔表达式 1 {
   /* 在布尔表达式 1 为 true 时执行 */
   if 布尔表达式 2 {
      /* 在布尔表达式 2 为 true 时执行 */
   }
}
```

### switch 语句
```
switch var1 {
    case val1:
        ...
    case val2, val3:
        ...
    default:
        ...
}
```

### Type Switch
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
```
switch x.(type){
    case type1:
       statement1    
    case type2:
       statement2 
    default: /* 可选 */
       statement3;
}
```

### select语句
select 语句类似于switch语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

这个东西我还真没有在其他语言里面看到过（可能是我了解的太少）。
```
select {
    case communication clause  :
       statement1;      
    case communication clause  :
       statement2; 
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement3;
}
```

- 每个case都必须是一个通信
- 所有channel表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，它就执行；其他被忽略。
- 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。 
- 否则：如果有default子句，则执行该语句。如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

学习这个需要先学习channel。我channel还不会呢，先不深入看这个了。


此系列其他文章地址
----------
https://github.com/zhangqunshi/golang_study


