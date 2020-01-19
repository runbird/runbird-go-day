package main

import (
	"errors"
	"fmt"
)

//A.
//var s []int
//s = append(s,1)

//B.
//var m map[string]int
//m["one"] = 1

//C.
//var s []int
//s = make([]int, 0)
//s = append(s,1)

//D.
//var m map[string]int
//m = make(map[string]int)
//m["one"] = 1
//ACD  B map may be nil

var ErrDidNotWork = errors.New("did't work")

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}
func second_120(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

//都输出 nil。知识点：变量的作用域。因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量。
func fix_second_120(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func main() {
	fmt.Println(second_120(true))
	fmt.Println(second_120(true))
	fmt.Println(fix_second_120(true))
	fmt.Println(fix_second_120(false))
}
