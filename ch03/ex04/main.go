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
		width,
		height,
		100,
		30.0,
		math.Pi / 6,
	}

	c.scale.xy = float64(c.width / 2 / c.xyrange)
	c.scale.z = float64(c.height) * 0.4

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: #%s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < c.cells; i++ {
		for j := 0; j < c.cells; j++ {
			ax, ay := c.corner(i+1, j)
			bx, by := c.corner(i, j)
			cx, cy := c.corner(i, j+1)
			dx, dy := c.corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

//!-
