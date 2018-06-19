# reading notes

## [dup1](dup1/main.go)

1. `if` : As with for, parentheses are never used around the condition in an `if` statement, but braces are required for the body.
2. `map` :
  - A set of key/value pairs , provides constant-time operations to store , retrieve
    > The key may be of any type whose values can compared with ==, strings being the most common example; the value may be of any type at all.
  - `make` creates a new empty map
  - It's not a problem if the map doesn't yet contain that key. `map[key]`` evaluates to the zero value for its type
  - The order of map iteration is not specified, but in practice it is random, varying from one run to another.This design is intentional, since it prevents programs from relying on any particular ordering where none is guaranteed.
3. `bufio` package, which helps make input and output efficient and convenient.
4. `fmt.Printf` : formatted output, like printf in C

## [dup2](dup2/main.go)

1. `file` : The function `os.Open` returns two values. The first is an open file (*os.File) that is used in subsequent reads by the Scanner.
2. `error` : built-in error type. If `err` equals the special built-in value `nil`, the file was opened successfully. On the other hand, if `err` is not nil, something went wrong. In that case, the error value describes the problem.
3. Functions and other package-level entities may be declared in any order
4. A `map` is a reference to the data structure create d by make.
  > When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller’s map reference too.

## [dup3](dup3/main.go)

1. `ReadFile` : `ioutil.ReadFile` reads the entire contents of a named file
2. `strings.Split` : splits a string into a slice of substrings

## [ex1.4](ex1.4/main.go)

1. `struct`
2. `bytes.Buffer`

# 读书笔记

## [dup1](dup1/main.go)

1. `if` : 语句条件两边不加括号，但是主体部分需要加
2. `map` :
  - key/value 键值对集合, 存、取均为常数时间操作
    > key 要能使用==运算符比较，value可以为任意类型
  - `make` 使用make创建空map
  - map中不含有的key，`map[key]`将得到其类型的0值
  - map的迭代顺序不确定。有意设计，防止程序依赖特定的遍历顺序
3. `bufio` 处理输入和输出方便又高效
4. `fmt.Printf` : 格式化字符串输出

## [dup2](dup2/main.go)

1. `file` : 使用`os.Open`操作文件
2. `error` : 内置错误类型. 等于内置空值 `nil` 表示没有错误,否则需要处理
3. 函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用
4. `map`是一个由make函数创建的数据结构的引用
  > map作为参数传递给某函数时，该函数接收这个引用的一份copy，被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到。

## [dup3](dup3/main.go)

1. `ReadFile` : `ioutil.ReadFile` 读取指定文件的全部内容
2. `strings.Split` : 把字符串分割成子串的切片

## [ex1.4](ex1.4/main.go)

1. `struct` : 结构体
2. `bytes.Buffer` : 高效处理字符串拼接
