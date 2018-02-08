Go语言学习笔记一
==============
听说Go语言又快又简单。即具有C语言的运行速度，又具有Python语言的开发效率，不知道真的假的。所以特意来学学这门“老”语言。

### 下载Go

先从简单的Hello world开始吧。首先从官网（https://golang.org/）下载Go。
我下载了windows版本的go1.9.4.windows-amd64.msi文件（91M）。公司网速好快，大约花了3秒就下载完成（实际上我连下载过程都没有看到）。
顺便说一下，go语言的吉祥物(mascot)是土拨鼠（Gopher），官方首页上就有，但是感觉画的真不咋地。

双击下载的文件就安装了，只要一直选择下一步（next）就行了。非常easy。

### 配置环境变量
现在只要是个语言，都需要配置环境变量，主要是把C:\Go\bin目录配置追加到PATH变量中（在系统--》高级--》环境变量中）。
不要问我怎么配，自己百度一下。


### 开始写Hello World程序
使用Notepad++建立一个文件hello.go （扩展名为go）：
```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

### 编译程序
通过cmd打开一个dos窗口。进入你文件所在目录：
```
cd d:\golang\src

go build
```

这样应该在src目录下面生成一个src.exe文件。这里比较让我奇怪的是，编译参数中没有输入要编译的文件名，而且生成好的文件名叫做src.exe。应该是直接使用当前所在目录的名称。不过看起来确实简单不少。

### 运行文件
在dos窗口直接输入`src.exe`回车就会显示hello, world。

看上去确实简单，不过有些地方感觉比较怪。Printf函数为啥要首字母大写呢。package main是一个特殊的包，表示其为可执行文件，不是一个库（都已经有main函数了，还整出这个概念干啥）。

-------------------------
参考资料：
https://golang.org/doc/

