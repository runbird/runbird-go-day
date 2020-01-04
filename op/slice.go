//https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466739&idx=4&sn=9de0eba9bdb36aa3b67d8a3903c53584&chksm=f247435cc530ca4a98abd9c1b2d58a63c7716e64236076c1296ed2964cd4fb68f7c8a6d44f85&scene=21#wechat_redirect
package main

import "fmt"

//nil切片: 这段代码声明了一个nil切片s，其实，切片的零值就是nil
func slice_1() {
	var s []int
	fmt.Println(s == nil)       // true
	fmt.Println(len(s), cap(s)) //0 0
	//切片就是一个数组的引用。
	//切片的类型在初始化时已经确认，就是[]Type，上面的代码就声明了[]int类型的nil切片s。
	//nil切片的指向底层数组的指针为nil。
}

//空切片：
func slice_2() {
	// 1、使用 make 创建空的整型切片
	s := make([]int, 0)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	// 2、使用切片字面量创建空的整型切片
	s2 := []int{}
	fmt.Println(s2)               // 输出：[]
	fmt.Println(len(s2), cap(s2)) // 输出：0 0
	//与nil切片一样，空切片的长度和容量也都是 0，说明切片底层的数组大小为 0，
	//是一个空数组（没有分配任何的存储空间）。
}

//声明了nil切片s1和三个非空切片s2、s3和s4。
//因为s1是nil切片，执行完copy操作之后，s1依然还是nil。这有别于append函数：
func slice_copy() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}
	// 1、
	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("s1 == nil", s1 == nil)
	// 2、
	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3) //s2: 4 5 6
	// 3、
	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4) //s3: 1 2 3 7
}

//由于s2的长度是 3，s3的长度是 4，所以执行copy操作只会从s3复制 3 个元素至s2。copy只会复制，不会追加。
func slice_append() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...)
	fmt.Println(s1) // 输出：[1 2 3]
}

//函数间传递切片：
//切片在函数间以值的方式传递。由于切片的尺寸很小（在 64 位架构的机器上，一个切片需要 24 字节的内存：
//指针字段、长度和容量字段各需要 8 字节），在函数间复制和传递切片成本也很低。
//切片发生复制时，底层数组不会被复制，数组大小也不会有影响。
func modify(s []int) {
	fmt.Printf("%p\n", &s)
	s[1] = 10
}

//不管是使用 nil 切片还是空切片，对其调用内置函数append、len和cap的效果都是一样的。
func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &s)
	modify(s)
	fmt.Println(s)
	//原切片地址和传递之后的切片的地址是不一样的，说明发生了复制；在函数modify中修改了切片一个值，原切片的值也随之改变了，说明这两个切片是共享底层数组的。
	//在函数间传递切片非常高效，而且不需要传递指针和处理复杂的语法，只需要复制切片，按自己的业务修改数据，最后传递回一个新的切片副本即可，这也是为什么函数间使用切片传参，而不是数组传参的原因。
}

//删除切片中的元素: Go没有提供删除切片元素的函数
func slice_delete() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s[:2], s[3:]...) // 删除索引为2的元素
	fmt.Println(s)              //[1 2 4 5 6]
}
