# reading notes

- Go natively handles Unicode, so it can process text in all the world’s languages.
- Go code is organized into packages. A package consists of one or more .go source files in a single directory that define what the package does.
  - **package** : package declaration is that states which package the file belongs to.
  - **import** :  other packages that it imports. You must import exactly the packages you need. A program will not compile if there are missing imports or if there are unnecessary ones.
  - *package main* : defines a standalone executable program.
  - *function main* : it's where execution of the program begins.
- Go does not require semicolons at the ends of statements or declarations, except where two or more appear on the same line.
  > In effect, newlines following certain tokens are converted into semicolons, so where newlines are placed matters to proper parsing of Go code. For instance, the opening brace { of the function must be on the same line as the end of the func declaration, not on a line by itself, and in the expression x + y, a newline is permitted after but not before the + operator.
- **gofmt** : The `gomft` tool rewrites code into the standard format.
- **goimports** : `goimports` additionally manages the insertion and removal of `import` declarations as needed.

# 读书笔记

- Go语言原生支持Unicode，它可以处理全世界任何语言的文本。
- Go语言的代码通过包（package）组织。一个包由位于单个目录下的一个或多个.go源代码文件组成, 目录定义包的作用。
  - **package** : 包（package）声明语句表示文件属于哪个包.
  - **import** : import语句导入其他包.必须恰当导入需要的包，缺少了必要的包或者导入了不需要的包，程序都无法编译通过。
  - *package main* : 定义了一个独立可执行的程序.
  -  *function main* : 整个程序执行时的入口
- Go语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。
 > 实际上，编译器会主动把特定符号后的换行符转换为分号, 因此换行符添加的位置会影响Go代码的正确解析（译注：比如行末是标识符、整数、浮点数、虚数、字符或字符串文字、关键字break、continue、fallthrough或return中的一个、运算符和分隔符++、--、)、]或}中的一个）。举个例子, 函数的左括号{必须和func函数声明在同一行上, 且位于末尾，不能独占一行，而在表达式x + y中，可在+后换行，不能在+前换行（译注：以+结尾的话不会被插入分号分隔符，但是以x结尾的话则会被分号分隔符，从而导致编译错误）。
- **gofmt** : `gofmt`工具把代码格式化为标准格式
- **goimports** : `goimports`可以根据代码需要, 自动地添加或删除`import`声明
