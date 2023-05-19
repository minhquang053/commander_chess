package cmdc

type Piece interface {
	canMove(board Board, start Spot, end Spot) bool
	getSide() bool
}

func (b *Board) takePiece(spot Spot) {

}
