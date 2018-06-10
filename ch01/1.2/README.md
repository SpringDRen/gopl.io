# reading notes

- `os.Args` : The first element of os.Args, os.Args[0], is the name of the command itself; the other elements are the arguments that were presented to the program when it started execution.
  > The os package provides functions and other values for dealing with the operating system in a platform-independent fashion. Command-line arguments are available to a program in a variable named Args that is part of the os package; thus its name any where outside the os package is os.Args.
- `slice` : a dynamically sized sequences of array elements where individual elements can be accessed as s[i] and a contiguous subsequence as s[m:n].Go uses half-open intervals that include the first index but exclude the last.
- variable declaration :
  - `s := ""`  //a short variable declaration, is the most compact, but it may be used only within a function, not for package-level variables.
  - `var s string`  //relies on default initialization to the zero value for strings, which is ""
  - `var s = ""`  //rarely used except when declaring multiple variables.
  - `var s string = ""`  //Use when variable’s type is not the same as initial value
- `i++` or `i--` : These are statements, not expressions. So j = i++ is illegal, and they are postfix only, so --i is not legal either.
- `for` : The `for` loop is the only loop statement in Go. Parentheses are never used around the three components of a for loop. The braces are mandatory, however, and the opening brace must be on the same line as the post statement.
```go
for initialization; condition; post {
    // zero or more statements
}
```
- `range` : In each iteration of the loop, `range` produces a pair of values: the index and the value of the element at that index.
- `_` : discard unused local variables
- `strings.Join` : A simpler and more efficient solution for concat-string

# 读书笔记

- `os.Args` : os.Args的第一个元素，os.Args[0], 是命令本身的名字；其它的元素则是程序启动时传给它的参数。  
 > os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。  
- `slice` : 切片，简版的动态数组。用s[i]访问单个元素，用s[m:n]获取子序列(左闭右开)
- variable declaration : 变量声明
  - `s := ""`  //短变量声明，最简洁，但只能用在函数内部，而不能用于包变量
  - `var s string`  //依赖于字符串的默认初始化零值机制，被初始化为""
  - `var s = ""`  //用得很少，除非同时声明多个变量
  - `var s string = ""`  //显式地标明变量的类型，用于当变量类型与初值类型不同时
- `i++` or `i--`: 它们是语句，而不是表达式。所以j = i++非法，而且++和--都只能放在变量名后面
- `for` : Go语言只有for循环这一种循环语句。for循环三个部分不需括号包围。大括号强制要求, 左大括号必须和post语句在同一行。
```go
for initialization; condition; post {
    // zero or more statements
}
```
- `range` : 在某种数据类型的区间（range）上遍历。`range`产生一对值；索引以及在该索引处的元素值。
- `_` : 丢弃无用的局部变量
- `strings.Join` : 字符串连接的高效解决方案
