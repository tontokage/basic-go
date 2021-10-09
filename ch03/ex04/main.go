// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"example.com/basic-go/ch03/ex04/canvas"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	q := r.URL.Query()
	height, _ := strconv.Atoi(q.Get("height"))
	width, _ := strconv.Atoi(q.Get("width"))
	color := q.Get("color")
	c := canvas.Canvas{
		Width:   width,
		Height:  height,
		Cells:   100.0,
		Xyrange: 30.0,
		Angle:   math.Pi / 6,
		Color:   color,
	}

	c.Scale.Xy = float64(float64(c.Width) / float64(2) / c.Xyrange)
	c.Scale.Z = float64(c.Height) * 0.4

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < c.Cells; i++ {
		for j := 0; j < c.Cells; j++ {
			ax, ay := c.Corner(i+1, j)
			bx, by := c.Corner(i, j)
			cx, cy := c.Corner(i, j+1)
			dx, dy := c.Corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'  fill='#%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.Color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

//!-
