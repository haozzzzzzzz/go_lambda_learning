package main

import (
	"fmt"
)

func main() {
	var a []string
	if a == nil {
		fmt.Println("nil")
	}

	a = make([]string, 0)
	if a == nil {
		fmt.Println("nil 2")
	}
}
