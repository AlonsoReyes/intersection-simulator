package car_generic

import (
	"fmt"
	m "github.com/AlonsoReyes/intersection-simulator/vehicle"
	"math"
)

func (car *Car) GetPosition() m.Pos {
	return car.Position
}

func (car *Car) Run(dt float64) {
	turnAngle := GetTurnAngle()
	car.Accelerate(dt)
	car.Forward(dt)
	car.Turn(dt, turnAngle)
}

func (car *Car) Forward(dt float64) {
	rad := car.Direction * (math.Pi / 180)
	posXDiff := math.Cos(rad) * car.Speed * dt // delta time
	posYDiff := math.Sin(rad) * car.Speed * dt
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
	// See the proportion between the covered distance and the initial distance until entering and exiting the danger zone
	startPos := GetEnterPosition(car.Lane, car.CoopZoneLength, car.DangerZoneLength)
	endPos := GetEndPosition(car.Lane, car.Intention, car.CoopZoneLength, car.DangerZoneLength)
	curPos := car.GetPosition()

	initDistance := startPos.GetManhattanDistance(endPos)
	curDistance := curPos.GetManhattanDistance(endPos)
	radio := GetTurnRadius(car.Intention, car.DangerZoneLength)
	// Why is it times 3?
	beautifulVariable := (turnAngle - car.ChangedAngle) * (1 - (curDistance / initDistance))
	dirChange := math.Abs(beautifulVariable * dt * car.Speed * (math.Pi / (2 * radio)))
	if car.ChangedAngle < turnAngle {
		//fmt.Println(car.ChangedAngle)
		car.ChangedAngle += dirChange
		if car.Intention == LeftIntention {
			car.Direction += 2.5 * dirChange
		} else {
			car.Direction -= dirChange
		}
	}
	// TODO check case when direction ends up negative a < destination < b (caso en que mi lane final es la de abajo)
	if car.Direction < 0 {
		car.Direction = 360 + car.Direction
	} else {
		car.Direction = math.Mod(car.Direction, 360)
	}
}

func checkTurnCondition(car *Car) bool {
	A, B, C, D := GetDangerZoneCoords(car.DangerZoneLength, car.CoopZoneLength)
	return IsInsideDangerZone(A, B, C, D, car.Position)
}

// TODO
func (car *Car) Turn(dt, turnAngle float64) {
	// Check if its in position to turn or not
	// check intention
	if car.Intention != StraightIntention {
		if checkTurnCondition(car) {
			car.ChangeDirection(dt, turnAngle)
		} else {
			if 0 < car.ChangedAngle && car.ChangedAngle != turnAngle {
				fmt.Println("LOOOL")
				car.ChangedAngle = turnAngle
				if car.Intention == LeftIntention {
					car.Direction = GetStartDirection(car.Lane) + turnAngle
				} else {
					car.Direction = GetStartDirection(car.Lane) - turnAngle
				}
			}
			fmt.Println(car.Position)
		}
	}
}
