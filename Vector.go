package vectors

import (
	"math"
)

type Vector2 interface {
	IntVal() (int, int)
	FloatVal() (float64, float64)
	Add(v Vector2) Vector2
	Sub(v Vector2) Vector2
	Muli(s int) Vector2
	Mulf(s float64) Vector2
	Abs() Vector2
	Normalise() Vector2f
	Magnitude() float64
	Right() Vector2
	Left() Vector2
	Back() Vector2
}

func IntVector(v Vector2) Vector2i {
	x, y := v.IntVal()

	vn := Vector2i{
		X: x,
		Y: y,
	}

	return vn
}

func FloatVector(v Vector2) Vector2f {
	x, y := v.FloatVal()

	vn := Vector2f{
		X: x,
		Y: y,
	}

	return vn
}

func Distance(v1 Vector2, v2 Vector2) float64 {
	x1, y1 := v1.FloatVal()
	x2, y2 := v2.FloatVal()

	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
