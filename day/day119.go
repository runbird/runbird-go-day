package main

import "fmt"

//1、下面说法正确的是。
//
//A. Go 语言中，声明的常量未使用会报错；
//B. cap() 函数适用于 array、slice、map 和 channel;
//C. 空指针解析会触发异常；
//D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值；
//cd

const (
	_        = iota
	c119 int = (10 * iota)
	d119
	e119 = iota
)

func second_119() {
	fmt.Printf("%d - %d - %d", c119, d119, e119)
}

func main() {
	second_119()
}
