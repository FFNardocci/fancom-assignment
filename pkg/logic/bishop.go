package logic

import "math"

func CanBishopAttack(bishopCoordinates, targetCoordinates []int) bool {
	return math.Abs(float64(bishopCoordinates[0]-targetCoordinates[0])) == math.Abs(float64(bishopCoordinates[1]-targetCoordinates[1]))
}
