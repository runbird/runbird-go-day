package main

//https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466739&idx=3&sn=9a1c9599172a532297ef41238450f9af&chksm=f247435cc530ca4ac6d92bd22011b52ae34d25d6e0eddf5b2ba85a15d3846b6674aa8bcf1d07&scene=21#wechat_redirect
import "fmt"

//因为原切片的容量已经满了，执行 append 操作之后会创建一个新的底层数组，
//并将原切片底层数组的值拷贝到新的数组，原数组保持不变。
func first_113() {
	x := []int{100, 200, 300, 400, 500, 600, 700}
	twohundred := &x[1]
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred)
}

func first_113_2() {
	x := make([]int, 0, 7)
	x = append(x, 100, 200, 300, 400, 500, 600, 700)
	twohundred := &x[1]
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred)

	x = make([]int, 0, 8) //指向一个新的切片
	x = append(x, 100, 200, 300, 400, 500, 600, 700)
	twohundred = &x[1]
	x = append(x, 800) //append 容量够，不会申请新的内存
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred)
}

//对一个切片执行 [i,j] 的时候，i 和 j 都不能超过切片的长度值
func second_113() {
	a := []int{0, 1}
	fmt.Printf("%v", a[len(a):])
}

func main() {
	first_113()
	second_113()
}
