package vehicle

import "math"

type Pos struct {
	X, Y float64
}

// Get angle with x as its center
func GetInsideAngle(p, x, v Pos) float64 {
	a := math.Pow(x.GetVectorLength(p), 2) + math.Pow(x.GetVectorLength(v), 2) - math.Pow(p.GetVectorLength(v), 2)
	b := 2.0 * x.GetVectorLength(p) * x.GetVectorLength(v)
	return math.Acos(a/b) * 180.0 / math.Pi
}

func (p *Pos) GetManhattanDistance(v Pos) float64 {
	return math.Abs(v.X-p.X) + math.Abs(v.Y-p.Y)
}

func (p *Pos) GetVectorLength(v Pos) float64 {
	return math.Sqrt(math.Pow(v.X-p.X, 2) + math.Pow(v.Y-p.Y, 2))
}

func (p *Pos) GetVector(v Pos) Pos {
	return Pos{v.X - p.X, v.Y - p.Y}
}

func (p *Pos) ScalarProduct(v Pos) float64 {
	return p.X*v.X + p.Y*v.Y
}

type Vehicle interface {
	GetPosition() Pos
	Run(dt float64)
	Forward(dt float64)
	Accelerate(dt float64)
	ChangeDirection(dt, turnAngle float64)
}
