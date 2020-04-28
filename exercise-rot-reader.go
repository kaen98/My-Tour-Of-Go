package main

import (
	"io"
	"os"
	"strings"
)

const src string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const des string = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(bytes []byte) (int, error) {
	len, err := rot13.r.Read(bytes)
	for i, v := range bytes {
		index := strings.IndexByte(src, v)
		if (index != -1) {
			bytes[i] = des[index]
		}
	}
	return len, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}