package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

type numInt int

func CalcSquare(sideLen float64, sidesNum numInt) float64 {
	var res float64
	switch sidesNum {
	case 0:
		res = math.Pi * math.Pow(sideLen, 2)
	case 3:
		res = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case 4:
		res = math.Pow(sideLen, 2)
	}
	return res
}
