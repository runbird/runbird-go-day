package main

import "fmt"

const (
	one = 1 << iota
	two
)

func first_117() {
	fmt.Println(one, two)
}

const (
	greeting = "Hello, Go"
	one1     = 1 << iota
	two1
)

func second_117() {
	fmt.Println(one1, two1)
}

func main() {
	first_117()
	second_117()
}
