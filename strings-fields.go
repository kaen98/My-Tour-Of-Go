package main

import (
	"fmt"
	"strings"
)

func main() {
	var fields = strings.Fields("  foo bar  baz   ")
	fmt.Printf("Type: %T, Fields are: %q", fields, fields)
}