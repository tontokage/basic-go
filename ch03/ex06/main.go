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
		y1 := (float64(py)-0.5)/height*(ymax-ymin) + ymin
		y2 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := (float64(px)-0.5)/width*(xmax-xmin) + xmin
			x2 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			m1 := mandelbrot(complex(x1, y1))
			m2 := mandelbrot(complex(x1, y2))
			m3 := mandelbrot(complex(x2, y1))
			m4 := mandelbrot(complex(x2, y2))
			avg := (m1 + m2 + m3 + m4) / 4
			// Image point (px, py) represents complex value z.
			img.Set(px, py, color.Gray{uint8(avg)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) int {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0
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
