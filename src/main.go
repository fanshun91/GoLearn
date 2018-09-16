package main

import (
	"fmt"
	"mypkg"
	"myapp"
	"calc"
	"calc/multi"
	_ "alias"
)

var myApp = myapp.PkgName() 

var multiVal = multi.Multi(43, 4)

func init() {
	fmt.Print("\n在main函数执行之前调用.")
}

func main() {
	mypkg.MypkgFunc()
	fmt.Print("\n", myApp)

	result := calc.Add(2, 3)
	fmt.Print("\n", "addVal: ", result)

	fmt.Print("\n", "multiVal: ", multiVal)

	fmt.Print("\n", "Hello World!")
}