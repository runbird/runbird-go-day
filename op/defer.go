package main

import "fmt"

//什么是 defer
//defer 是 Go 语言提供的一种用于注册延迟调用的机制，每一次 defer 都会把函数压入栈中，
//当前函数返回前再把延迟函数取出并执行。

//defer 语句并不会马上执行，而是会进入一个栈，函数 return 前，会按先进后出（FILO）的顺序执行。
//也就是说最先被定义的 defer 语句最后执行。先进后出的原因是后面定义的函数可能会依赖前面的资源，自然要先执行；
//否则，如果前面先执行，那后面函数的依赖就没有了。
func main() {
	fmt.Println(increaseB())
	fmt.Println(f1(), f2(), f3())
}

//采坑点
//使用 defer 最容易采坑的地方是和带命名返回参数的函数一起使用时。
//defer 语句定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
//作为函数参数，则在 defer 定义时就把值传递给 defer，并被缓存起来；作为闭包引用的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值。
//避免掉坑的关键是要理解这条语句：
//return xxx这条语句并不是一个原子指令，经过编译之后，变成了三条指令：
//
//1. 返回值 = xxx
//2. 调用 defer 函数
//3. 空的 return
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
	//实际

	// 1.赋值
	//r = 0
	// 2.闭包引用，返回值被修改
	//defer func() {
	//	r++
	//}()
	// 3.空的 return
	//return
	//defer 是闭包引用，返回值被修改，所以 f() 返回 1
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
	//t := 5
	// 1.赋值
	//r = t
	// 2.闭包引用，但是没有修改返回值 r
	//defer func() {
	//	t = t + 5
	// 若为 r += 5 -> return结果为10
	//}()
	// 3.空的 return
	//return
	//第二步没涉及返回值 r 的操作，所以返回 5。
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1

	// 1.赋值
	//r = 1
	// 2.r 作为函数参数，不会修改要返回的那个 r 值
	//defer func(r int) {
	//	r = r + 5
	//}(r)
	// 3.空的 return
	//return
	//第二步，r 是作为函数参数使用，是一份复制，defer 语句里面的 r 和 外面的 r 其实是两个变量，里面变量的改变不会影响外层变量 r，所以不是返回 6 ，而是返回 1。
}
