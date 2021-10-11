package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(bitDiffCount(&c1, &c2))
}

func bitDiffCount(c1, c2 *[32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		a := c1[i] ^ c2[i]
		count += int(pc[a])
	}
	return count
}
