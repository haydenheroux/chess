package main

type Board map[Square]Piece

func NewBoard() Board {
	board := make(Board)

	board.Sets("a8", NewRook(Black))
	board.Sets("h8", NewRook(Black))

	board.Sets("b8", NewKnight(Black))
	board.Sets("g8", NewKnight(Black))

	board.Sets("c8", NewBishop(Black))
	board.Sets("f8", NewBishop(Black))

	board.Sets("d8", NewQueen(Black))
	board.Sets("e8", NewKing(Black))

	for file := 0; file < 8; file++ {
		square := Square{Rank: 6, File: file}
		board.Set(square, NewPawn(Black))
	}

	for file := 0; file < 8; file++ {
		square := Square{Rank: 1, File: file}
		board.Set(square, NewPawn(White))
	}

	board.Sets("a1", NewRook(White))
	board.Sets("h1", NewRook(White))

	board.Sets("b1", NewKnight(White))
	board.Sets("g1", NewKnight(White))

	board.Sets("c1", NewBishop(White))
	board.Sets("f1", NewBishop(White))

	board.Sets("d1", NewQueen(White))
	board.Sets("e1", NewKing(White))

	return board
}

func (board Board) Set(square Square, piece Piece) bool {
	if empty := board.IsEmpty(square); empty {
		board[square] = piece
		return true
	}
	return false
}

func (board Board) Sets(s string, piece Piece) bool {
	return board.Set(Notation(s), piece)
}

func (board Board) Get(square Square) (Piece, bool) {
	val, ok := board[square]
	return val, ok
}

func (board Board) Remove(square Square) {
	delete(board, square)
}

func (board Board) Pop(square Square) (Piece, bool) {
	if piece, occupied := board.Get(square); occupied {
		board.Remove(square)
		return piece, true
	}

	return Piece{}, false
}

func (board Board) IsNotEmpty(square Square) bool {
	_, isOccupied := board.Get(square)
	return isOccupied
}

func (board Board) IsEmpty(square Square) bool {
	return !board.IsNotEmpty(square)
}

func (board Board) Move(from Square, to Square) bool {
	fromIsEmpty := board.IsEmpty(from)
	toIsNotEmpty := board.IsNotEmpty(to)
	if fromIsEmpty || toIsNotEmpty {
		return false
	}

	piece, _ := board.Pop(from)
	board[to] = piece
	return true
}

func (board Board) IsAlly(square Square, side Side) bool {
	piece, isOccupied := board.Get(square)
	if isOccupied {
		return piece.Side == side
	}
	return false
}

func (board Board) IsEnemy(square Square, side Side) bool {
	piece, isOccupied := board.Get(square)
	if isOccupied {
		return piece.Side != side
	}
	return false
}
