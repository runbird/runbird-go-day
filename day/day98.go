package main

import "fmt"

func first_98() {
	//A. map 反序列化时 json.unmarshal() 的入参必须为map的地址；
	//
	//B. 在函数调用中传递 map，则子函数中对 map 元素的增加不会导致父函数中 map 的修改；
	//
	//C. 在函数调用中传递 map，则子函数中对 map 元素的修改不会导致父函数中 map 的修改；
	//
	//D. 不能使用内置函数 delete() 删除 map 的元素；

	//op/gomap.go
}

//------------
type Foo struct {
	val int
}

func (f Foo) Inc(inc int) {
	f.val += inc
}

//使用值类型接收者定义的方法，调用的时候，使用的是值的副本，对副本操作不会影响的原来的值。
// 如果想要在调用函数中修改原值，可以使用指针接收者定义的方法。
func second_98() {
	var f Foo
	f.Inc(100)
	fmt.Println(f.val)
}

func (f *Foo) Inc2(inc int) {
	f.val += inc
}

func fix_second_98() {
	f := &Foo{}
	f.Inc2(100)
	fmt.Println(f.val)
}

func main() {
	second_98()
	fix_second_98()
}
