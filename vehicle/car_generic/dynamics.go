package car_generic

import (
	m "github.com/niclabs/intersection-simulator/vehicle"
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

/*
Changes the direction toward a 90 degree turn so it heads to the corresponding exit lane
*/
func (car *Car) ChangeDirection(dt, turnAngle float64) {
	// See the proportion between the covered distance and the initial distance until entering and exiting the danger zone
	startPos := GetEnterPosition(car.Lane, car.CoopZoneLength, car.DangerZoneLength)
	endPos := GetEndPosition(car.Lane, car.Intention, car.CoopZoneLength, car.DangerZoneLength)
	curPos := car.GetPosition()

	turnCenterPos := GetCenterOfTurn(startPos, endPos, car)
	coveredAngle := m.GetInsideAngle(startPos, turnCenterPos, curPos)

	dirChange := math.Abs(coveredAngle - car.ChangedAngle)
	if car.ChangedAngle < turnAngle {
		car.ChangedAngle += dirChange
		if car.Intention == LeftIntention {
			car.Direction += dirChange
		} else {
			car.Direction -= dirChange
		}
	}

	if car.Direction < 0 {
		car.Direction = 360 + car.Direction
	} else {
		car.Direction = math.Mod(car.Direction, 360)
	}
}

/*
Checks if the car is inside the intersection. Possible collision zone.
*/
func checkTurnCondition(car *Car) bool {
	A, B, C, D := GetDangerZoneCoords(car.DangerZoneLength, car.CoopZoneLength)
	return IsInsideDangerZone(A, B, C, D, car.Position)
}

/*
Checks if the car needs to turn depending of what amount of the turnAngle it has covered.
In a four-way intersection this angle is 90 degrees.
*/
func (car *Car) Turn(dt, turnAngle float64) {
	if car.Intention != StraightIntention {
		if checkTurnCondition(car) {
			car.ChangeDirection(dt, turnAngle)
		} else {
			if 0 < car.ChangedAngle && car.ChangedAngle != turnAngle {
				car.ChangedAngle = turnAngle
				if car.Intention == LeftIntention {
					car.Direction = GetStartDirection(car.Lane) + turnAngle
				} else {
					car.Direction = GetStartDirection(car.Lane) - turnAngle
				}
			}
		}
	}
}
