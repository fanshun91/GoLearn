package vars

// import "fmt"
import (
	"fmt"
	"reflect"
)

var (
	a int = 12
	b string = "name"
)

// var n1, n2 = 1, "str"

// 全局变量必须使用var声明
// var num = 1

func init() {
	// var e, f = 1, 2
	//冒号简写的方式只能用在函数体内
	e, f := 1, 2

	// 类型转换
	var x = 3
	y := float32(x)

	// fmt.Println("\n::", x)

	fmt.Print("\n", a, "\n", b)
	fmt.Print("\n", reflect.TypeOf(b))
	fmt.Print("\n", e, "\n", f)
	fmt.Print("\n", y)
	fmt.Print("\n", reflect.TypeOf(y))
}