package main

import (
	"fmt"
	"time"
)

type foo struct {
	Val int
}

type bar struct {
	Val int
}

//go无引用变量，每个变量都占用一个唯一的内存位置，所以第一个比较 false
func first_96() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Println(c)
	fmt.Println(foo(d))
	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(f)
	fmt.Print(a == b, c == foo(d), e == f)
}

func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}
func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}
//1 和 2随机输出
func second_96() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)
}

func main() {
	first_96()
	fmt.Println()
	second_96()
}
