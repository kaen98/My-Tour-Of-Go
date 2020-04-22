package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dx)
	for x, _ := range pic {
		pic[x] = make([]uint8, dy) 
	}

	for x, v := range pic {
		for y, _ := range v {
			pic[x][y] = uint8(image_func_c(x, y))
		} 
	}

	return pic
}

func image_func_a(x, y int) uint8 {
	return uint8((x + y) / 2);
}

func image_func_b(x, y int) uint8 {
	return uint8(x * y);
}

func image_func_c(x, y int) uint8 {
	return uint8(x ^ y);
}



func main() {
	pic.Show(Pic)
}