package main

type Board map[Square]Piece

func NewBoard() Board {
	board := make(Board)

	board[Notation("a8")] = NewRook(Notation("a8"), Black)
	board[Notation("h8")] = NewRook(Notation("h8"), Black)

	board[Notation("b8")] = NewKnight(Notation("b8"), Black)
	board[Notation("g8")] = NewKnight(Notation("g8"), Black)

	board[Notation("c8")] = NewBishop(Notation("c8"), Black)
	board[Notation("f8")] = NewBishop(Notation("f8"), Black)

	board[Notation("d8")] = NewQueen(Notation("d8"), Black)
	board[Notation("e8")] = NewKing(Notation("e8"), Black)

	for file := 0; file < 8; file++ {
		square := Square{Rank: 6, File: file}
		board[square] = NewPawn(square, Black)
	}

	for file := 0; file < 8; file++ {
		square := Square{Rank: 1, File: file}
		board[square] = NewPawn(square, White)
	}

	board[Notation("a1")] = NewRook(Notation("a1"), White)
	board[Notation("h1")] = NewRook(Notation("h1"), White)

	board[Notation("b1")] = NewKnight(Notation("b1"), White)
	board[Notation("g1")] = NewKnight(Notation("g1"), White)

	board[Notation("c1")] = NewBishop(Notation("c1"), White)
	board[Notation("f1")] = NewBishop(Notation("f1"), White)

	board[Notation("d1")] = NewQueen(Notation("d1"), White)
	board[Notation("e1")] = NewKing(Notation("e1"), White)

	return board
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
	piece.Square = to
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
