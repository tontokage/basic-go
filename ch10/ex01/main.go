// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt" // register PNG decoder
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	format := flag.String("format", "jpeg", "output image format")
	flag.Parse()
	img, err := decode(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch *format {
	case "jpeg":
		jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95})
	case "png":
		png.Encode(os.Stdout, img)
	case "gif":
		gif.Encode(os.Stdout, img, &gif.Options{})
	}
}

func decode(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
