package main

import (
	"fmt"
)

// 传参使用的是值拷贝，可能会造成性能问题，通常建议使用slice

func test(x [2]int) {
    fmt.Printf("x: %p\n", &x)
    x[1] = 1000
}

func main() {
    a := [2]int{}
    fmt.Printf("a: %p\n", &a)

    test(a)
	fmt.Println(a)
	// 内置函数 len 和 cap 都返回数组长度 (元素数量)。
	fmt.Println(len(a),cap(a))
}