package cmdc

import "math"

type Commander struct {
	Notation string
	Side     bool
}

func (c Commander) getSide() bool {
	return c.Side
}

func (c Commander) canMove(board Board, start Spot, end Spot) bool {
	if start.X != end.X && start.Y != end.Y {
		return false
	}
	xDiff := int(math.Abs(float64(start.X - end.X)))
	yDiff := int(math.Abs(float64(start.Y - end.Y)))
	if xDiff > 10 || yDiff > 10 {
		return false
	}

	if piece, exist := board.Spots[end]; exist {
		if c.getSide() == piece.getSide() {
			return false
		}
		if xDiff+yDiff != 1 {
			return false
		}
	}

	for spot, piece := range board.Spots {
		if _, isCommander := piece.(Commander); isCommander && spot.X == end.X {
			return false
		}
		if (spot.X < end.X && spot.X > start.X && spot.Y == start.Y) ||
			(spot.X > end.X && spot.X < start.X && spot.Y == start.Y) ||
			(spot.Y < end.Y && spot.Y > start.Y && spot.X == start.X) ||
			(spot.Y > end.Y && spot.Y < start.Y && spot.X == start.X) {
			return false
		}
	}

	return true
}
