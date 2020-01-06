package main

import (
	"fmt"
	"time"
)

//CD
func first_102() {
	//关于循环语句，下面说法正确的有？
	//
	//A. 循环语句既支持 for 关键字，也支持 while 和 do-while；
	//B. 关键字for的基本使用方法与C/C++中没有任何差异；
	//C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；
	//D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；
}

//var ch chan int = make(chan int)
var ch = make(chan int)

func generate_102() {
	for i := 17; i < 5000; i += 17 {
		ch <- i
		time.Sleep(1 * time.Millisecond)
	}
	close(ch)
}

func second_102() {
	timeout := time.After(800 * time.Millisecond)
	go generate_102()
	found := 0
	for {
		select {
		case i, ok := <-ch:
			if ok {
				if i%38 == 0 {
					fmt.Println(i, "is a multiple of 17 and 38")
					found++
					//break 会跳出 select 块，但不会跳出 for 循环。
					if found == 3 {
						break
					}
				}
			} else {
				break
			}
		case <-timeout:
			fmt.Println("timed out")
			break
		}
	}
	//unreachable code ： 因为不能跳出for循环，不会执行
	fmt.Println("The end")
}

//
func fix_second_102() {
	timeout := time.After(800 * time.Millisecond)
	go generate_102()
	found := 0
MAIN_LOOP:
	for {
		select {
		case i, ok := <-ch:
			if ok {
				if i%38 == 0 {
					fmt.Println(i, "is a multiple of 17 and 38")
					found++
					if found == 3 {
						break MAIN_LOOP
					}
				}
			} else {
				break MAIN_LOOP
			}
		case <-timeout:
			fmt.Println("timed out")
			break MAIN_LOOP
		}
	}
	fmt.Println("The end")
}

func main() {
	second_102()
}
