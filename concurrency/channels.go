package main

import "fmt"

func sum(s int, c chan int) {
	c <- s // send sum to c
}

func main() {
	c := make(chan int, 1)
	go sum(17, c)
	go sum(5, c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}