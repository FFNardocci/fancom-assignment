package logic

func CanBishopAttack(bishopCoordinates, targetCoordinates []int) bool {
	return (bishopCoordinates[0]+bishopCoordinates[1] == targetCoordinates[0]+targetCoordinates[1]) ||
		(bishopCoordinates[0]-bishopCoordinates[1] == targetCoordinates[0]-targetCoordinates[1])
}
