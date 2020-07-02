package main


import (
	"fmt"
	"math"
)


func main()  {
	// 1.\" 转义
	fmt.Println("str := \"c:\\pprof\\main.exe\"")

	// 2.多行字符串
	s1 := `第一行
	第二行
	第三行
	`
	fmt.Println(s1)

	// 3.byte和rune类型 
	// 字符串是一个byte数组
	// 一个rune代表一个字符串，byte（unint8）类型代表一个字节
	traversalString()

	// 4.修改字符串
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	changeString()

	// 5.类型转换
	// Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。强制类型转换的基本语法如下：
	// T(表达式)
	sqrtDemo()
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}


func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}