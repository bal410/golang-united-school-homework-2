package square

import "math"

type sidesInt uint8

const (
	SidesCircle   sidesInt = 0
	SidesTriangle sidesInt = 3
	SidesSquare   sidesInt = 4
)

func CalcSquare(sideLen float64, sidesNum sidesInt) float64 {
	return math.Pow(sideLen, 2) * getSquareCoef(sidesNum)
}

func getSquareCoef(sidesNum sidesInt) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi
	case SidesTriangle:
		return math.Sqrt(3) / 4
	case SidesSquare:
		return 1
	default:
		return 0
	}
}
