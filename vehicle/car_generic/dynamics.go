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

	initDistance := startPos.GetVectorLength(endPos)
	curDistance := curPos.GetVectorLength(endPos)

	dirChange := turnAngle * (1 - (curDistance / initDistance)) * dt
	/* after
	radio := GetTurnRadius(car.Intention, car.DangerZoneLength)
	dirChange := turnAngle * car.Speed * dt * (math.Pi / (2 * radio))
	*/
	fmt.Println(startPos, endPos)
	//fmt.Println((curDistance / initDistance), dirChange, car.ChangedAngle)
	if car.ChangedAngle < 90 {
		car.ChangedAngle += dirChange
		if car.Intention == LeftIntention {
			car.Direction += dirChange
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
	if car.Intention != StraightIntention {
		A, B, C, D := GetDangerZoneCoords(car.DangerZoneLength, car.CoopZoneLength)
		return IsInsideDangerZone(A, B, C, D, car.Position)
	}
	return false
}

// TODO
func (car *Car) Turn(dt, turnAngle float64) {
	// Check if its in position to turn or not
	// check intention
	if checkTurnCondition(car) {
		fmt.Println("qeeee")
		car.ChangeDirection(dt, turnAngle)
	}
}
