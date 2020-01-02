package main

import "fmt"

type Point struct {
	x, y int
}

//for range循环的时候，获取到的元素值是副本
func first_97() {
	points := []Point{
		{1, 2},
		{3, 4},}

	for _, p := range points {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(points)
}

func fix_first_97() {
	points := []*Point{
		&Point{1, 2},
		&Point{3, 4},}
	for _, p := range points {
		p.x, p.y = p.y, p.x
	}
	fmt.Println("{} ,{}", points[0], points[1])
}

//get()返回的切片和原切片公用底层数组，如果在调用函数里面修改返回的切片，将会影响到原切片
func get() []byte {
	raw := make([]byte, 1000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

func second_97() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])
}

func fix_get() []byte{
	raw := make([]byte, 1000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw)
	return res;
}

func fix_second_97(){
	data := fix_get()
	fmt.Println(len(data), cap(data), &data[0])
}
func main() {
	first_97()
	fix_first_97()
	second_97()
	fix_second_97()

	//空阻塞
	select {

	}
}
