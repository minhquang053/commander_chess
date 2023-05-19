package cmdc

type Board struct {
	Spots map[Spot]Piece
}

type Spot struct {
	X int
	Y int
}
