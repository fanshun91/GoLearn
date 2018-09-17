package constv

import "fmt"

const a = iota // 0
const b = iota // 0

// 插值使用法
const (
	c = iota // 0
	d = iota // 1
	_ // 2
	e = iota // 3
	f = 32
	g = iota // 4
)

// 隐式使用法
const (
	h = iota * 2 // 0
	i // `1*2` 2
	j // `2*2` 4
)

const (
	x, y = iota, iota * 3 // 0, 0
	z, o // 1, 3
)