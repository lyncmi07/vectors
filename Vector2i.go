package vectors;

import (
  "math"
)

type Vector2i struct {
  X int;
  Y int;
}

func (v Vector2i) IntVal() (int, int) {
  return v.X, v.Y;
}
func (v Vector2i) FloatVal() (float64, float64) {
  return float64(v.X), float64(v.Y);
}

func (vo Vector2i) Add(v Vector2) Vector2 {
  x, y := v.IntVal();
  vo.X += x;
  vo.Y += y;

  return vo;
}

func (vo Vector2i) Sub(v Vector2) Vector2 {
  x, y := v.IntVal();
  vo.X -= x;
  vo.Y -= y;

  return vo;
}

func (v Vector2i) Abs() Vector2 {
  v.X = int(math.Abs(float64(v.X)));
  v.Y = int(math.Abs(float64(v.Y)));

  return v;
}

func (v Vector2i) Muli(s int) Vector2 {
  v.X *= s;
  v.Y *= s;

  return v;
}

func (v Vector2i) Mulf(s float64) Vector2 {
  vn := Vector2f{X:float64(v.X), Y:float64(v.Y)};

  vn.X *= s;
  vn.Y *= s;

  return vn;
}

func (v Vector2i) Normalise() Vector2f {
  x := math.Abs(float64(v.X));
  y := math.Abs(float64(v.Y));

  xn := float64(v.X) / (x + y);
  yn := float64(v.Y) / (x + y);

  vn := Vector2f{X:xn, Y:yn};

  return vn;
}

func (v Vector2i) Magnitude() float64 {
  x := math.Abs(float64(v.X));
  y := math.Abs(float64(v.Y));

  return x + y;
}

func (v Vector2i) Left() Vector2 {
  v.X = v.Y;
  v.Y = -v.X;

  return v;
}

func (v Vector2i) Right() Vector2 {
  v.X = -v.Y;
  v.Y = v.X;

  return v
}

func (v Vector2i) Back() Vector2 {
  v.X = -v.X;
  v.Y = -v.Y;

  return v;
}
