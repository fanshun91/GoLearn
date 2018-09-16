package main

import (
	"fmt"
	// "time"
)

// 常量声明
const NAME = "golang"

// 变量声明
var globalVar string = "variables"

// 结构声明
type Learn struct {}

// 声明接口 
type Ilearn interface {}

// 一般变量声明
type imoocInt int

// 在main函数外部定义自定义函数
func learnGo() {
	fmt.Print("learn go!! \n")
}

// main函数 用于生成可执行文件
func main() {
	// learnGo()
	fmt.Print("hello world \n")
	// time.Sleep(3*time.Second)
	fmt.Print(globalVar)
}