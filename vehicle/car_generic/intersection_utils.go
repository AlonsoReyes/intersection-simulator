package car_generic

import (
	"fmt"
	m "github.com/AlonsoReyes/intersection-simulator/vehicle"
)

func GetTurnAngle() float64 {
	return FourwayTurnAngle
}

func GetTurnRadius(intention int, dangerZoneLength float64) float64 {
	var radius float64
	if intention == LeftIntention {
		radius = dangerZoneLength * (3 / 4)
	} else if intention == RightIntention {
		radius = dangerZoneLength / 4
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
		x = (coopZoneLength + laneWidth) / 2
		y = 0
	case RightLane:
		x = coopZoneLength
		y = (coopZoneLength + laneWidth) / 2
	case TopLane:
		x = (coopZoneLength - laneWidth) / 2
		y = coopZoneLength
	case LeftLane:
		x = 0
		y = (coopZoneLength - laneWidth) / 2
	default:
		x = 0
		y = 0
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
	fmt.Println(A, B, C, D, M)
	am := A.ScalarProduct(M)
	ab := A.ScalarProduct(B)
	ad := A.ScalarProduct(D)
	return 0 <  (am * ab) && (am * ab) < (ab * ab) && 0 <  (am * ad) && (am * ad) < (ad * ad)
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