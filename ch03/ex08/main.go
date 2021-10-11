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
	"math/big"
	"math/cmplx"
	"os"
)

const (
	windowSize    = .00001
	xmin, ymin    = -.5430, -.6157
	xmax, ymax    = xmin + windowSize, ymin + windowSize
	width, height = 1024, 1024
)

func main() {
	//img := generateImgComplex64()
	//img := generateImgComplex128()
	img := generateImgComplexBigFloat()
	png.Encode(os.Stdout, img)
}

func generateImgComplex64() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrotComplex64(complex64(z)))
		}
	}
	return img
}

func mandelbrotComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func generateImgComplex128() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}

	return img
}

func mandelbrotComplex128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

type ComplexBigFloat struct {
	real, imag *big.Float
}

func (c *ComplexBigFloat) Add(a, b *ComplexBigFloat) *ComplexBigFloat {
	c.real = new(big.Float).Add(a.real, b.real)
	c.imag = new(big.Float).Add(a.imag, b.imag)
	return c
}

//  (a + bi) * (c + di)
// =ac + adi + bci + bdi^2
// = ab - bd + i(ad + bc)
func (c *ComplexBigFloat) Mul(a, b *ComplexBigFloat) *ComplexBigFloat {
	c.real = new(big.Float).Sub(
		new(big.Float).Mul(a.real, b.real),
		new(big.Float).Mul(a.imag, b.imag),
	)
	c.imag = new(big.Float).Add(
		new(big.Float).Mul(a.real, b.imag),
		new(big.Float).Mul(b.real, a.imag),
	)
	return c
}

func (c *ComplexBigFloat) Abs(a *ComplexBigFloat) *big.Float {
	return new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(a.real, a.real),
			new(big.Float).Mul(a.imag, a.imag),
		),
	)
}

func generateImgComplexBigFloat() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		//y := big.NewFloat(float64(py)/height*(ymax-ymin) + ymin)
		y := new(big.Float).Add(
			new(big.Float).Mul(
				new(big.Float).Quo(big.NewFloat(float64(py)), big.NewFloat(height)),
				big.NewFloat(ymax-ymin),
			),
			big.NewFloat(ymin),
		)
		for px := 0; px < width; px++ {
			//x := big.NewFloat(float64(px)/width*(xmax-xmin) + xmin)
			x := new(big.Float).Add(
				new(big.Float).Mul(
					new(big.Float).Quo(big.NewFloat(float64(px)), big.NewFloat(width)),
					big.NewFloat(xmax-xmin),
				),
				big.NewFloat(xmin),
			)
			z := &ComplexBigFloat{x, y}
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrotComplexBigFloat(z))
		}
	}
	return img

}

func mandelbrotComplexBigFloat(z *ComplexBigFloat) color.Color {
	const iterations = 200
	const contrast = 15

	v := &ComplexBigFloat{big.NewFloat(0), big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		//v = v*v + z
		v = new(ComplexBigFloat).Add(
			new(ComplexBigFloat).Mul(v, v),
			z,
		)
		if new(ComplexBigFloat).Abs(v).Cmp(big.NewFloat(2)) > 0 {
			return color.Gray{255 - contrast*n}
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
