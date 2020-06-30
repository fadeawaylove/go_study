package main

import "fmt"

/*
var 变量名 变量类型
*/

func foo() (int, string) {
	return 10, "Q1mi"
}

func main() {
	// 1.批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	a = "123"
	b = 2
	c = true
	d = 3.2
	fmt.Println(a, b, c, d)

	// 2.变量的初始化 var 变量名 类型 = 表达式
	var name string = "pprof.cn"
	var sex int = 1
	// 一次初始化多个变量，省略类型编译器可以根据右边的值来推导变量类型
	//  var name, sex = "pprof.cn", 1
	fmt.Println(name, sex)

	// 3. :=声明变量
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)

	// 4.匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	// 常量
	// const pi = 3.1415
	// const e = 2.7182
	const (
		pi = 3.1415
		e  = 2.7182
	)
	fmt.Println(pi, e)
	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	// iota iota是go语言的常量计数器，只能在常量的表达式中使用

}
