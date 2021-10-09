package canvas

import (
	"math"
)

type Scale struct {
	Xy float64
	Z  float64
}

type Canvas struct {
	Width   int
	Height  int
	Cells   int
	Xyrange float64
	Color   string
	Scale   Scale
	Angle   float64
}

func (c Canvas) Corner(i, j int) (float64, float64) {
	var sin30, cos30 = math.Sin(c.Angle), math.Cos(c.Angle) // sin(30°), cos(30°)
	// Find point (x,y) at corner of cell (i,j).
	x := c.Xyrange * (float64(i)/float64(c.Cells) - 0.5)
	y := c.Xyrange * (float64(j)/float64(c.Cells) - 0.5)

	// Compute surface height z.
	z := c.f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(c.Width)/2 + (x-y)*cos30*c.Scale.Xy
	sy := float64(c.Height)/2 + (x+y)*sin30*c.Scale.Xy - z*c.Scale.Z
	return sx, sy
}

func (c Canvas) f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
