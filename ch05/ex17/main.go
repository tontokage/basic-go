package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	for _, n := range ElementByTagName(doc, "img", "div") {
		fmt.Printf("<%s>\n", n.Data)
	}
}

func ElementByTagName(n *html.Node, names ...string) []*html.Node {
	nodes := []*html.Node{}
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				nodes = append(nodes, n)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		children := ElementByTagName(c, names...)
		nodes = append(nodes, children...)
	}

	return nodes
}
