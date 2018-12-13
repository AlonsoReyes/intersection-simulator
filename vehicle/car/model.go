package car

import (
	"fmt"
	"math"
	dyn "github.com/AlonsoReyes/intersection-simulator/vehicle/dynamics"
)

type Car struct {
	Position                            dyn.Pos
	Intention, Lane                int // Left = 0, Right = 1
	Direction, Acceleration, Speed float64
}

func CreateCar(x, y, speed float64, intention, lane int) *Car {
	pos := dyn.Pos{x, y}
	car := Car{Position: pos, Intention: intention, Lane: lane, Speed: speed}
	return &car
}

func (car *Car) PrintCar() {
	fmt.Println(car.Position.X)
	fmt.Println(car.Position.Y)
	fmt.Println(car.Intention)
	fmt.Println(car.Lane)
	fmt.Println(car.Direction)
	fmt.Println(car.Speed)
	fmt.Println(car.Acceleration)
}

func (car *Car) GetPosition() (float64, float64){
	return car.Position.X, car.Position.Y
}

func (car *Car) Move(dt float64) {
	car.Accelerate(dt)
	car.Advance(dt)
	car.Turn(dt, TurnAngle)
}

func (car *Car) Advance(dt float64) {
	rad := car.Direction * (math.Pi / 180)
	posXDiff := math.Sin(rad) * car.Speed * dt // delta time
	posYDiff := math.Cos(rad) * car.Speed * dt
	car.Position.X += posXDiff
	car.Position.Y += posYDiff
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

// Changes the direction toward a 90 degree turn towards the corresponding direction of the intention
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
	if checkTurnCondition(car) || car.Intention == StraightIntention  {
		return
	}
	car.ChangeDirection(dt, turnAngle)
}

func checkTurnCondition(car *Car) bool {

}
