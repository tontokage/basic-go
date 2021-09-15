package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	execTime := execCase1(os.Args, start)
	execTime2 := execCase2(os.Args, start)

	fmt.Println(execTime - execTime2)
}

// case 1 inefficiency way
func execCase1(args []string, t time.Time) time.Duration {
	s, sep := "", " "
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	return time.Since(t)
}

// case 2 strings.Join
func execCase2(args []string, t time.Time) time.Duration {
	start := time.Now()
	s := strings.Join(os.Args, " ")
	fmt.Println(s)
	return time.Since(start)
}
