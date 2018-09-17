package process

import "fmt"

var a = [] string {"a", "b", "c"}

func init() {
  for key, val := range a {
		fmt.Print("\nkey: ", key, ", val: ", val)
	}

	// 如果不想使用key
	// for _, val := range a { ... }

	// goto的使用
	goto One
	fmt.Print("中间代码块")
	One: 
		fmt.Print("这里是代码块一")
		
	// 用goto做无限循环
	// One: 
	// 	fmt.Print("这里是代码块一")
	// goto One
}
