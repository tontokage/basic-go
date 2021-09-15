// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/basic-go/ch02/ex02/length"

	"gopl.io/ch2/tempconv"
)

func main() {
	var input []string
	if len(os.Args) > 1 {
		input = os.Args[1:]
	} else {
		fmt.Print("数値を入力してください")
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			fmt.Print("入力エラー")
			os.Exit(1)
		}
		input = append(input, in)
	}
	for _, arg := range input {
		printTemp(arg)
		printLength(arg)
	}
}

//!-
func printTemp(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func printLength(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := length.Feet(t)
	m := length.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, length.FToM(f), m, length.MToF(m))
}
