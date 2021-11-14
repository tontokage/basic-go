// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(u string) []string {
	fmt.Println(u)
	hostName, err := getHostName(u)
	if err != nil {
		log.Print(err)
	}
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	refined := refineByHostName(list, hostName)
	fmt.Println(refined)
	return refined
}

func getHostName(u string) (string, error) {
	ur, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return ur.Host, nil
}

func refineByHostName(list []string, h string) []string {
	// 多分非効率
	ret := []string{}
	for _, v := range list {
		vHost, err := getHostName(v)
		if err != nil {
			log.Print(err)
		}
		if vHost == h {
			ret = append(ret, v)
		}
	}
	return ret
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
