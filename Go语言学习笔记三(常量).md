Go语言学习笔记三： 常量
==============

### 定义常量
常量就是在声明后不能再修改的量。
```
const x int = 100
const y string = "abc"
const z = "abc"
```
看上去与变量的定义差不多，就是把var变成了const关键字。

## 枚举
```
const (
    Yellow = 1
    Red = 2
    Blue = 3
)
```

### 特殊常量iota
这个iota非常奇葩，从0开始，每次使用一次就自动增加一。例如：
```
const (
    a = iota
    b = iota
    c = iota
)
```
上面a=0，b=1，c=2。不知道为什么要发明这么奇怪的东西，难道是因为懒，不过你需要至少用4次以上，否则就不划算了。


此系列其他文章地址
----------
https://github.com/zhangqunshi/golang_study


