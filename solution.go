package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type Sides int

const (
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
	SidesCircle   Sides = 0
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SidesTriangle:
		return sideLen * math.Sqrt(3.0) / 4.0
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen * sideLen * math.Pi / 2.0
	default:
		return 0.0
	}
}
