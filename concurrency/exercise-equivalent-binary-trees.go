package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func goWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	var t1_list []int
	var t2_list []int


	go goWalk(t1, ch1)
	go goWalk(t2, ch2)

	for v1 := range ch1 {
		t1_list = append(t1_list, v1)
	}

	for v2 := range ch2 {
		t2_list = append(t2_list, v2)
	}

	if (len(t1_list) != len(t2_list)) {
		return false
	}



	for i, _ := range t1_list {
		if (t1_list[i] != t2_list[i]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
}
