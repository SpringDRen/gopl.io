package main

import "fmt"

var Version string = "unset"

//编译参数 -ldflags
//go run -ldflags '-X main.Version=1.0' main.go
//go build -ldflags '-X main.Version=1.0' main.go

//linux 环境下还可以执行命令。e.g.:go run -ldflags "-X 'main.VersionString=`date`'" main.go
//Version: Sun Nov 27 16:42:10 EST 2016
func main() {
	fmt.Printf("Version:%s\n", Version)
}
