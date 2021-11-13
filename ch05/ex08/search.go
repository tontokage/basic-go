// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%#v", ElementByID(doc, "header"))
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var target *html.Node
	search := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					target = n
					return false
				}
			}

		}
		return true
	}

	//!+call
	forEachNode(doc, search, nil)
	//!-call

	return target
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if n == nil {
		return
	}

	if pre != nil {
		if next := pre(n); !next {
			return
		}
	}

	forEachNode(n.FirstChild, pre, post)
	forEachNode(n.NextSibling, pre, post)
	if post != nil {
		post(n)
	}
}
