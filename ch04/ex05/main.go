package main

import "fmt"

func main() {
	s := []string{"a", "a", "b"}
	uniq(s)
	fmt.Println(s)
}

func uniq(s []string) []string {
	if len(s) < 2 {
		return s
	}
	j := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			s[j] = s[i]
			j++
		}
	}
	return s[:j]
}
