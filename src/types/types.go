package types

import (
	"fmt"
	"unsafe"
)

var uint8i uint8 = 1

var int32i int32 = 1

var intdef int = 1

var floati float32 = 1.3

var booli bool = false

var runei rune = 1

// 通过type来定义数据类型的别名
// type 整型 int32

func init() {
	fmt.Print(
		"\n",
		"uint8: ",
		unsafe.Sizeof(uint8i),
		"\n",
		"int32: ",
		unsafe.Sizeof(int32i),
		"\n",
		"int: ",
		unsafe.Sizeof(intdef),
	)
}