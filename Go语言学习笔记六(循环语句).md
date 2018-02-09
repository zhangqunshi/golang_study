Go语言学习笔记六： 循环语句
==============

今天学了一个格式化代码的命令：`gofmt -w chapter6.go`

### for循环
for循环有3种形式：
```
for init; condition; increment {
}

// 类似while
for condition {
}

// 和for(;;)一样
for {
}
```
循环slice，map，数组，字符串还可以使用下面这种方式：
```
for key, value := range oldMap {
    newMap[key] = value
}
```

### break 语句
- 用于循环语句中跳出循环，并开始执行循环之后的语句。
- break在switch（开关语句）中在执行一条case后跳出语句的作用。

### continue 语句
- Go 语言的 continue 语句有点像 break 语句。但是continue不是跳出循环，而是跳过当前循环执行下一次循环语句。
- for 循环中，执行 continue 语句会触发for增量语句的执行。

### goto 语句

- Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
- goto语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
```
goto label;
...
...
label: statement;
```

为什么又要把goto搞进来，难道设计者就是喜欢折磨我们。


此系列其他文章地址
----------
https://github.com/zhangqunshi/golang_study


