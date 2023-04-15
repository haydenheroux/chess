package main

func (piece Piece) IsLegalMove(square Square) bool {
	for _, move := range piece.LegalMoves(square) {
		if move == square {
			return true
		}
	}

	return false
}

func (piece Piece) LegalMoves(square Square) []Square {
	return []Square{}
}
