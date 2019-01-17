package car_generic

import (
	m "github.com/niclabs/intersection-simulator/vehicle"
)

const (
	MaxSpeed = 200.0
	MinSpeed = 0.0

	// Turn intentions
	LeftIntention     = 0
	RightIntention    = 1
	StraightIntention = 2

	// Lanes
	BottomLane = 0
	RightLane  = 1
	TopLane    = 2
	LeftLane   = 3

	// How much the direction must change when turning, this is just for a square fourway
	FourwayTurnAngle = 90.0
)

// Lengths will be used from the config file in the intersection directory
type Car struct {
	ChangedAngle     float64 // Keeps track of how much the car has turned
	Position         m.Pos
	Intention        int
	Lane             int
	Direction        float64
	Acceleration     float64
	Speed            float64
	CoopZoneLength   float64 // Used to set the initial position of cars
	DangerZoneLength float64 // Inner intersection, if it is a square Height == Width, in the image used its 1/3 of the whole pic
}

func CreateCar(lane, intention int, coopZoneLength, dangerZoneLength float64) *Car {
	var c Car
	c.ChangedAngle = 0.0
	c.Position = GetStartPosition(lane, coopZoneLength, dangerZoneLength)
	c.Intention = intention
	c.Lane = lane
	c.Direction = GetStartDirection(lane)
	c.Acceleration = 5
	c.Speed = 100
	c.CoopZoneLength = coopZoneLength
	c.DangerZoneLength = dangerZoneLength
	return &c
}
