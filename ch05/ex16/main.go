package main

import (
	"errors"
	"fmt"
)

func main() {
	s, _ := join(",", "a", "b", "c", "", "d")
	fmt.Printf("%s\n", s)
}

func join(sep string, ss ...string) (string, error) {
	if len(ss) == 0 {
		return "", errors.New("func join need at least 1 string")
	}

	ret := string(ss[0])
	for _, s := range ss[1:] {
		ret += sep + string(s)
	}
	return ret, nil
}
