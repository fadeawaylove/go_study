package main

import (
	"fmt"
)


/*
需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

    1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。 
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/



func main() {
	// 1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	}else{
		fmt.Println("不是空")
	}
	// 2. :=
	s2 := []int{}
	fmt.Println(s2)
	// 3. make
	var s3 []int = make([]int, 0)
	fmt.Println(s3)

	//4. 初始化赋值
	var s4 []int = make([]int, 3, 3)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	//5. 从数组切片
	arr := []int{1,2,3,4,5}
	var s6 []int
	s6 = arr[1:4]
	fmt.Println(s6)

}