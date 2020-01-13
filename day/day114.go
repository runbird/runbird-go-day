package main

import "fmt"

//complie error
//const (
//    ERR_ELEM_EXISTerror = errors.New("element already exists")
//    ERR_ELEM_NT_EXISTerror = errors.New("element not exists")
//)
func link(p ...interface{}) {
	fmt.Println(p)
}

//新切片
func main() {
	link("seek", 1, 2, 3, 4)
	a := []int{1, 2, 3, 4}
	link("seek", a)
	//修改下面的代码，使得第二个输出 [seek 1 2 3 4] 。

	templink := make([]interface{}, 0, len(a)+1)
	templink = append(templink, "seek")
	for _, i := range a {
		templink = append(templink, i)
	}
	link(templink...)
}
