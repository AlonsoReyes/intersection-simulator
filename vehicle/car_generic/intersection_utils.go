package car_generic

import (
	m "github.com/niclabs/intersection-simulator/vehicle"
	"math"
)

func GetTurnAngle() float64 {
	return FourwayTurnAngle
}

func GetTurnRadius(intention int, dangerZoneLength float64) float64 {
	var radius float64
	if intention == LeftIntention {
		radius = dangerZoneLength * (3.0 / 4.0)
	} else if intention == RightIntention {
		radius = dangerZoneLength / 4.0
	} else {
		radius = 0.0
	}
	return radius
}

func GetStartDirection(lane int) (dir float64) {
	switch lane {
	case BottomLane:
		dir = 90
	case RightLane:
		dir = 180
	case TopLane:
		dir = 270
	case LeftLane:
		dir = 0
	default:
		dir = 360
	}
	return dir
}

func GetStartPosition(lane int, coopZoneLength, dangerZoneLength float64) m.Pos {
	var x, y float64
	laneWidth := dangerZoneLength / 2.0
	switch lane {
	case BottomLane:
		x = (coopZoneLength + laneWidth) / 2.0
		y = 0
	case RightLane:
		x = coopZoneLength
		y = (coopZoneLength + laneWidth) / 2.0
	case TopLane:
		x = (coopZoneLength - laneWidth) / 2.0
		y = coopZoneLength
	case LeftLane:
		x = 0
		y = (coopZoneLength - laneWidth) / 2.0
	default:
		x = 0
		y = 0
	}
	res := m.Pos{X: x, Y: y}

	return res
}

/*
Returns the coordinates of where the car enters the inner intersection.
*/
func GetEnterPosition(lane int, coopZoneLength, dangerZoneLength float64) m.Pos {
	var x, y float64
	laneWidth := dangerZoneLength / 2.0
	switch lane {
	case BottomLane:
		x = (coopZoneLength + laneWidth) / 2.0
		y = (coopZoneLength - dangerZoneLength) / 2.0
	case TopLane:
		x = (coopZoneLength - laneWidth) / 2.0
		y = (coopZoneLength + dangerZoneLength) / 2.0
	case LeftLane:
		x = (coopZoneLength - dangerZoneLength) / 2.0
		y = (coopZoneLength + laneWidth) / 2.0
	case RightLane:
		x = (coopZoneLength + dangerZoneLength) / 2.0
		y = (coopZoneLength - laneWidth) / 2.0
	default:
		x = 0.0
		y = 0.0
	}
	res := m.Pos{X: x, Y: y}

	return res
}

/*
Returns the coordinates of where the car exits the inner intersection.
*/
func GetEndPosition(lane, intention int, coopZoneLength, dangerZoneLength float64) m.Pos {
	var x, y float64
	laneWidth := dangerZoneLength / 2.0
	switch lane {
	case BottomLane:
		if intention == LeftIntention {
			x = (coopZoneLength - dangerZoneLength) / 2.0
			y = (coopZoneLength + laneWidth) / 2.0
		} else if intention == RightIntention {
			x = (coopZoneLength + dangerZoneLength) / 2.0
			y = (coopZoneLength - laneWidth) / 2.0
		} else {
			x = (coopZoneLength + laneWidth) / 2.0
			y = (coopZoneLength + dangerZoneLength) / 2.0
		}
	case TopLane:
		if intention == LeftIntention {
			x = (coopZoneLength + dangerZoneLength) / 2.0
			y = (coopZoneLength - laneWidth) / 2.0
		} else if intention == RightIntention {
			x = (coopZoneLength - dangerZoneLength) / 2.0
			y = (coopZoneLength + laneWidth) / 2.0
		} else {
			x = (coopZoneLength - laneWidth) / 2.0
			y = (coopZoneLength - dangerZoneLength) / 2.0
		}
	case LeftLane:
		if intention == LeftIntention {
			x = (coopZoneLength + laneWidth) / 2.0
			y = (coopZoneLength + dangerZoneLength) / 2.0
		} else if intention == RightIntention {
			x = (coopZoneLength - laneWidth) / 2.0
			y = (coopZoneLength - dangerZoneLength) / 2.0
		} else {
			x = (coopZoneLength + dangerZoneLength) / 2.0
			y = (coopZoneLength - laneWidth) / 2.0
		}
	case RightLane:
		if intention == LeftIntention {
			x = (coopZoneLength - laneWidth) / 2.0
			y = (coopZoneLength - dangerZoneLength) / 2.0
		} else if intention == RightIntention {
			x = (coopZoneLength + laneWidth) / 2.0
			y = (coopZoneLength + dangerZoneLength) / 2.0
		} else {
			x = (coopZoneLength - dangerZoneLength) / 2.0
			y = (coopZoneLength + laneWidth) / 2.0
		}
	default:
		x = (coopZoneLength - dangerZoneLength) / 2.0
		y = (coopZoneLength + laneWidth) / 2.0
	}
	res := m.Pos{X: x, Y: y}

	return res
}

/*
	A		B
		M
	D		C

	Condition
	(0<AM⋅AB<AB⋅AB)∧(0<AM⋅AD<AD⋅AD)
*/

func IsInsideDangerZone(A, B, C, D, M m.Pos) bool {
	am := A.GetVector(M)
	ab := A.GetVector(B)
	ad := A.GetVector(D)
	AMAD := am.ScalarProduct(ad)
	AMAB := am.ScalarProduct(ab)
	ABAB := ab.ScalarProduct(ab)
	ADAD := ad.ScalarProduct(ad)
	return 0 < AMAB && AMAB < ABAB && 0 < AMAD && AMAD < ADAD
}

/*
	A		B

	D		C
*/

func GetDangerZoneCoords(dangerZoneLength, coopZoneLength float64) (m.Pos, m.Pos, m.Pos, m.Pos) {
	var A, B, C, D m.Pos
	laneWidth := dangerZoneLength / 2
	halfCoopZone := coopZoneLength / 2

	A.X = halfCoopZone - laneWidth
	A.Y = halfCoopZone + laneWidth

	B.X = halfCoopZone + laneWidth
	B.Y = halfCoopZone + laneWidth

	C.X = halfCoopZone + laneWidth
	C.Y = halfCoopZone - laneWidth

	D.X = halfCoopZone - laneWidth
	D.Y = halfCoopZone - laneWidth

	return A, B, C, D
}

/*
Returns the center of the circumference that represents the turn trajectory.
*/
func GetCenterOfTurn(startPos, endPos m.Pos, car *Car) m.Pos {
	startDir := GetStartDirection(car.Lane) * math.Pi / 180.0
	a := math.Abs(math.Cos(startDir))
	b := math.Abs(math.Sin(startDir))
	x := startPos.X*a + endPos.X*b
	y := startPos.Y*b + endPos.Y*a
	return m.Pos{X: x, Y: y}
}
