package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hoge$hoge moge$mogefuga"
	f := func(s string) string {
		return "#" + s
	}
	fmt.Println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	const start = "$"
	ret := ""
	targetS := s
	for {
		st := strings.Index(targetS, start)
		if st == -1 {
			ret += targetS
			break
		}
		ret += targetS[:st]
		en := strings.Index(s[st+1:], start)
		if en == -1 {
			en = len(targetS)
			ret += f(targetS[st+1 : en])
			break
		}
		ret += f(targetS[st+1 : en+1])
		targetS = targetS[en+1:]
	}

	return ret
}
