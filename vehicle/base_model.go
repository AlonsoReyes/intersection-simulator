package vehicle



type Pos struct {
	X, Y float64
}

func (p *Pos) ScalarProduct(v Pos) float64 {
	return p.X * v.X + p.Y * v.Y
}

type Vehicle interface {
	GetPosition() Pos
	Run(dt float64)
	Forward(dt float64)
	Accelerate(dt float64)
	ChangeDirection(dt,turnAngle float64)
}