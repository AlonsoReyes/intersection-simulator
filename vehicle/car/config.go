package car

import (
	dyn "github.com/AlonsoReyes/intersection-simulator/vehicle/dynamics"
)

/*
Lane distribution
	2

3		1

	0
*/

const (
	WinWidth = 768
	WinHeigth = 768
	CarWidth = WinWidth / 10
	CarHeight = WinHeigth / 10
	LaneWidth = CarWidth * 1.1

	// Intersection metrics
	IntersectionWidht = LaneWidth * 2
	IntersectionHeight = IntersectionWidht

	// Dynamics restrictions
	MaxSpeed = 40.0
	MinSpeed = 0.0
	MaxAcceleration = 6.0
	//  Turn variables
	LeftTurnRadio = IntersectionWidht / 4
	RightTurnRadio = IntersectionWidht * (3 / 4)

	TurnAngle = 90.0

	// Intentions
	LeftIntention = 0
	RightIntention = 1
	StraightIntention = 2

	// Lanes
	BottomLane = 0
	RightLane = 1
	TopLane = 2
	LeftLane = 3

	// Default initial values
	DefaultInitialSpeed = 0.0
	DefaultCarImage = "crashing-cars/images/redcar.png"
)


var Lanes = []dyn.Pos {
	dyn.Pos{(WinWidth + LaneWidth) / 2.0, 0 }, // 0
	dyn.Pos{WinWidth, (WinHeigth + LaneWidth) / 2.0 }, // 1
	dyn.Pos{(WinWidth - LaneWidth) / 2.0,  WinHeigth }, // 2
	dyn.Pos{0, (WinHeigth - LaneWidth) / 2.0 }, // 3
}

// The indexes represent the corresponding lanes
var InitialDirections = [] float64 {
	0.0,
	270.0,
	180.0,
	90.0,
}