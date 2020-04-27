package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}