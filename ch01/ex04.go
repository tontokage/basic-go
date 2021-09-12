// 重複した行のそれぞれが含まれていた全てのファイルの名前を表示するようにdup2を修正しなさい
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)

			f.Close()
		}

	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s fileName: %s\n", n, line, fileNames[line])
		}
	}

}

func countLines(f *os.File, counts map[string]int, fileNames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		l := input.Text()
		counts[l]++
		if fileNames[l] != f.Name() {
			fileNames[l] += f.Name()
		}
	}
}
