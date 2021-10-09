// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

//func mandelbrot(z complex128) color.Color {
//	const iterations = 200
//	const contrast = 15
//
//	var v complex128
//	for n := uint8(0); n < iterations; n++ {
//		v = v*v + z
//		if cmplx.Abs(v) > 2 {
//			return color.Gray{255 - contrast*n}
//		}
//	}
//	return color.Black
//}

func mandelbrot(z complex128) color.Color {
	const contrast = 85
	const iterations = 255 / contrast

	var v complex128
	var ncolor1, ncolor2 *uint8
	for i := 0; i <= 2; i++ {
		var nred, ngreen, nblue uint8
		switch i {
		case 0:
			ncolor1 = &nblue
			ncolor2 = &ngreen
		case 1:
			ncolor1 = &ngreen
			ncolor2 = &nred
		case 2:
			ncolor1 = &nred
			ncolor2 = &nblue
		}

		for n := uint8(0); n <= iterations; n++ {
			nc := n * contrast

			*ncolor1 = 255 - nc
			*ncolor2 = nc

			v = v*v + z
			if cmplx.Abs(v) > 2 {
				return color.RGBA{nred, ngreen, nblue, 255}
			}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
