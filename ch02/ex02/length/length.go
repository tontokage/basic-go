package length

import "fmt"

type Feet float64
type Meter float64

func FToM(f Feet) Meter { return Meter(f / 3.2808) }
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

func (f Feet) String() string  { return fmt.Sprintf("%gfeet", f) }
func (m Meter) String() string { return fmt.Sprintf("%gmeter", m) }
