package carmdl

import (
	"fmt"
	"math"
)

type Pos struct {
	X, Y float64
}

type Car struct {
	Pos                            Pos
	Intention, Lane                int // Left = 0, Right = 1
	Direction, Acceleration, Speed float64
}

func CreateCar(x, y, speed float64, intention, lane int) *Car {
	pos := Pos{x, y}
	car := Car{Pos: pos, Intention: intention, Lane: lane, Speed: speed}
	return &car
}

func (car *Car) PrintCar() {
	fmt.Println(car.Pos.X)
	fmt.Println(car.Pos.Y)
	fmt.Println(car.Intention)
	fmt.Println(car.Lane)
	fmt.Println(car.Direction)
	fmt.Println(car.Speed)
	fmt.Println(car.Acceleration)
}

func (car *Car) GetPosition() (float64, float64){
	return car.Pos.X, car.Pos.Y
}

func (car *Car) Move(dt float64) {
	car.Accelerate(dt)
	car.Move(dt)
	car.Turn(dt, TurnAngle)
}

func (car *Car) Advance(dt float64) {
	rad := car.Direction * (math.Pi / 180)
	posXDiff := math.Sin(rad) * car.Speed * dt// delta time
	posYDiff := math.Cos(rad) * car.Speed * dt
	car.Pos.X += posXDiff
	car.Pos.Y += posYDiff
}

func (car *Car) Accelerate(dt float64) {
	speedDiff := car.Acceleration * dt
	newSpeed := car.Speed + speedDiff
	if newSpeed > MaxSpeed {
		newSpeed = MaxSpeed
	} else if newSpeed < MinSpeed {
		newSpeed = MinSpeed
	}
	car.Speed = newSpeed
}

// Changes the direction toward a 90 degree turn towards the corresponding diretion of the intention
func (car *Car) ChangeDirection(dt, turnAngle float64) {
	radio := 0.0
	if car.Intention == LeftIntention {
		radio = LeftTurnRadio
	} else {
		radio = RightTurnRadio
	}
	dirChange := turnAngle * car.Speed * dt * (math.Pi / 2 * radio)
	if car.Intention == LeftIntention {
		car.Direction += dirChange
	} else {
		car.Direction -= dirChange
	}
}

// TODO
func (car *Car) Turn(dt, turnAngle float64) {
	// Check if its in position to turn or not
	if car.Intention == StraightIntention {
		return
	}
	if  {
		car.ChangeDirection(dt, turnAngle)
	}
}
