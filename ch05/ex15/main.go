package main

import (
	"errors"
	"fmt"
)

func main() {
	max, _ := max(0, 1, 2, 3, -1)
	min, _ := min(0, 1, 2, 3, -1)
	fmt.Printf("max is %v\n", max)
	fmt.Printf("min is %v\n", min)
}

func max(val ...int) (int, error) {
	if len(val) == 0 {
		return 0, errors.New("func max needs at least 1 number")
	}
	max := val[0]
	for _, v := range val {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func min(val ...int) (int, error) {
	if len(val) == 0 {
		return 0, errors.New("func max needs at least 1 number")
	}
	min := val[0]
	for _, v := range val {
		if v < min {
			min = v
		}
	}
	return min, nil
}
