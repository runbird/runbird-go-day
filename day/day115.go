package main

import (
	"fmt"
)

func first_115() {
	ns := []int{010: 200, 005: 100}
	fmt.Println(len(ns))
}

//如果不注释的话ch <- f()向通道中塞了一个值，通道就满了，select里面就塞不进去了，所在走的default，你要不把通道里面的值消耗掉，要么把通道容量设置大一点。
//
//1.消耗值，在 ch<-f() 后面加上<-ch取出一个值，现在通道又空了；
//2.将ch := make(chan int, 1)这条语句里面1改为2
func second_115() {
	i := 0
	f := func() int {
		i++
		return i
	}
	ch := make(chan int, 1)
	//	ch <- f()
	select {
	case ch <- f():
		fmt.Println("hello")
	default:
		fmt.Println(i)
	}
}

func main() {
	first_115()
	second_115()
}
