package main

import (
	"bufio"
	"fmt"
)

func main() {
	var lc LineCounter
	name := "shun"
	fmt.Fprintf(&lc, "hello, %s.\nI'm Takaya.", name)
	fmt.Println(lc)

	var wc WordCounter
	name = "shun"
	fmt.Fprintf(&wc, "hello, %s.\nI'm Takaya.", name)
	fmt.Println(wc)
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			*c += 1
		}
	}
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanWords(p, true)

		if err != nil {
			return 0, err
		}
		if token != nil {
			*c++
		}

		p = p[advance:]
	}

	return n, nil
}
