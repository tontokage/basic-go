package canvas

import "math"

type Scale struct {
	xy float64
	z  float64
}

type Canvas struct {
	width   int
	height  int
	cells   int
	xyrange float64
	color   int
	scale   Scale
	angle   float64
}

func (c Canvas) Corner(i, j int) (float64, float64) {
	var sin30, cos30 = math.Sin(c.angle), math.Cos(c.angle) // sin(30°), cos(30°)
	// Find point (x,y) at corner of cell (i,j).
	x := c.xyrange * (float64(i/c.cells) - 0.5)
	y := c.xyrange * (float64(j/c.cells) - 0.5)

	// Compute surface height z.
	z := c.f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(c.width)/2 + (x-y)*cos30*c.scale.xy
	sy := float64(c.height)/2 + (x+y)*sin30*c.scale.xy - z*c.scale.z
	return sx, sy
}

func (c Canvas) f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
