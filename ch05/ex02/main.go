// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	count := countElem(map[string]int{}, doc)
	for key, c := range count {
		fmt.Println(key, c)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func countElem(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count = countElem(count, c)
	}
	return count
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
