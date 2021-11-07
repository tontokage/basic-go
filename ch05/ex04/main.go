// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"errors"
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
	links := visit(map[string][]string{}, doc)
	for attr, linkList := range links {
		for _, link := range linkList {
			fmt.Println(attr, link)
		}
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links map[string][]string, n *html.Node) map[string][]string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			val, err := getAttr(n, "href")
			if err == nil {
				links[n.Data] = append(links[n.Data], val)
			}
		case "script", "style", "img":
			val, err := getAttr(n, "src")
			if err == nil {
				links[n.Data] = append(links[n.Data], val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func getAttr(n *html.Node, attr string) (string, error) {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val, nil
		}
	}
	return "", errors.New("no attribute " + attr)
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
