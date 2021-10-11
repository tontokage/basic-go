package main

import (
	"fmt"
	"os"
)

func main() {
	i1 := os.Args[1]
	i2 := os.Args[2]
	if isAnagram(i1, i2) {
		fmt.Printf("%sと%sはアナグラムです\n", i1, i2)
	} else {
		fmt.Printf("%sと%sはアナグラムではありません\n", i1, i2)
	}

}

func isAnagram(a, b string) bool {
	if a == "" || b == "" {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	runeMapA := generateRuneMap(a)
	runeMapB := generateRuneMap(b)

	for r, countA := range runeMapA {
		if countB, ok := runeMapB[r]; !ok && countA != countB {
			return false
		}
	}
	return true
}

func generateRuneMap(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		if r == ' ' {
			continue
		}
		count[r]++
	}
	return count
}

//func isAnagram(a, b string) bool {
//	main := bytes.NewBufferString(a)
//	target := bytes.NewBufferString(b)
//
//	if main.Len() != target.Len() {
//		return false
//	}
//
//	for i := 0; i < target.Len(); i++ {
//		target.
//	}
//
//	fmt.Println(main.String(), target.String())
//	return true
//}
