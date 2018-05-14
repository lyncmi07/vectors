package vectors;

import (
  "math"
)

type Vector2f struct {
  X float64;
  Y float64;
}

func (v Vector2f) IntVal() (int,int) {
  return int(v.X), int(v.Y);
}

func (v Vector2f) FloatVal() (float64,float64) {
  return v.X, v.Y
}

func (vo Vector2f) Add(v Vector2) Vector2 {
  x, y := v.FloatVal();

  vo.X += x;
  vo.Y += y;

  return vo;
}

func (vo Vector2f) Sub(v Vector2) Vector2 {
  x, y := v.FloatVal();

  vo.X -= x;
  vo.Y -= y;

  return vo;
}

func (v Vector2f) Abs() Vector2 {
  v.X = math.Abs(v.X);
  v.Y = math.Abs(v.Y);

  return v;
}

func (v Vector2f) Muli(s int) Vector2 {
  v.X *= float64(s);
  v.Y *= float64(s);

  return v;
}

func (v Vector2f) Mulf(s float64) Vector2 {
  v.X *= s;
  v.Y *= s;

  return v;
}

func (v Vector2f) Normalise() Vector2f {
  x := math.Abs(v.X);
  y := math.Abs(v.Y);

  xn := v.X / (x + y);
  yn := v.Y / (x + y);

  vn := Vector2f{X:xn, Y:yn};

  return vn;
}

func (v Vector2f) Magnitude() float64 {
  x := math.Abs(v.X);
  y := math.Abs(v.Y);

  return x + y;
}

func (v Vector2f) Left() Vector2 {
  v.X = v.Y
  v.Y = -v.X

  return v;
}

func (v Vector2f) Right() Vector2 {
  v.X = -v.Y
  v.Y = v.X

  return v;
}

func (v Vector2f) Back() Vector2 {
  v.X = -v.X;
  v.Y = -v.Y;

  return v;
}
