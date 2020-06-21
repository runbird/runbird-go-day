package main

import (
	"errors"
	"fmt"
)

func first_120() {
	var s []int
	s = append(s, 1)

	//error
	var m map[string]int
	m["one"] = 1

	var s2 []int
	s2 = make([]int, 0)
	s2 = append(s2, 1)

	var m2 map[string]int
	m2 = make(map[string]int)
	m2["one"] = 1
}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

//FIXED

//TODO
func DoTheThingFixed(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		// :=   --> =
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func second_120() {
	//都输出 nil。知识点：变量的作用域。因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量。
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThingFixed(true))

	//	fmt.Println(DoTheThing(true))
	//	fmt.Println(DoTheThing(false))
}

func main() {
	second_120()
}
