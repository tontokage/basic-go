package main

import (
	"fmt"

	"example.com/basic-go/ch02/ex03/popcount"
	//"gopl.io/ch2/popcount"
)

func main() {
	var x uint64 = 111111
	fmt.Println(popcount.PopCount(x))
}
