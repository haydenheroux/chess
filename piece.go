package main

import "fmt"

type Square struct {
	Rank int
	File int
}

func Notation(notation string) Square {
	rank := int(notation[1]-'0')-1
	file := int(notation[0]-'a')
	return Square{Rank: rank, File: file}
}

func (square Square) String() string {
	file := []string{"a", "b", "c", "d", "e", "f", "g", "h"}[square.File]
	return fmt.Sprintf("%s%d", file, square.Rank+1)
}

func (square Square) IsBlack() bool {
	return (square.Rank+square.File)%2 == 0
}

func (square Square) IsWhite() bool {
	return !square.IsBlack()
}

type PieceType int

const (
	Pawn PieceType = iota
	Bishop
	Knight
	Rook
	Queen
	King
)

func (pieceType PieceType) String() string {
	return []string{"p", "b", "k", "R", "Q", "K"}[pieceType]
}

type Side int

const (
	Black Side = iota
	White
)

func (side Side) String() string {
	return []string{"b", "w"}[side]
}

type Piece struct {
	Square Square
	Type   PieceType
	Side   Side
}

func NewPiece(pieceType PieceType, square Square, side Side) Piece {
	return Piece{
		Square: square,
		Type:   pieceType,
		Side:   side,
	}
}

func NewPawn(square Square, side Side) Piece {
	return NewPiece(Pawn, square, side)
}

func NewBishop(square Square, side Side) Piece {
	return NewPiece(Bishop, square, side)
}

func NewKnight(square Square, side Side) Piece {
	return NewPiece(Knight, square, side)
}

func NewRook(square Square, side Side) Piece {
	return NewPiece(Rook, square, side)
}

func NewQueen(square Square, side Side) Piece {
	return NewPiece(Queen, square, side)
}

func NewKing(square Square, side Side) Piece {
	return NewPiece(King, square, side)
}

type Board []Piece

func NewBoard() Board {
	var board Board

	board = append(board, NewRook(Notation("a8"), Black))
	board = append(board, NewRook(Notation("h8"), Black))

	board = append(board, NewKnight(Notation("b8"), Black))
	board = append(board, NewKnight(Notation("g8"), Black))

	board = append(board, NewBishop(Notation("c8"), Black))
	board = append(board, NewBishop(Notation("f8"), Black))

	board = append(board, NewQueen(Notation("d8"), Black))
	board = append(board, NewKing(Notation("e8"), Black))

	for file := 0; file < 8; file++ {
		square := Square{Rank: 6, File: file}
		board = append(board, NewPawn(square, Black))
	}

	for file := 0; file < 8; file++ {
		square := Square{Rank: 1, File: file}
		board = append(board, NewPawn(square, White))
	}

	board = append(board, NewRook(Notation("a1"), White))
	board = append(board, NewRook(Notation("h1"), White))

	board = append(board, NewKnight(Notation("b1"), White))
	board = append(board, NewKnight(Notation("g1"), White))

	board = append(board, NewBishop(Notation("c1"), White))
	board = append(board, NewBishop(Notation("f1"), White))

	board = append(board, NewQueen(Notation("d1"), White))
	board = append(board, NewKing(Notation("e1"), White))

	return board
}

func (board Board) Get(square Square) (Piece, bool) {
	for _, piece := range board {
		if piece.Square == square {
			return piece, true
		}
	}

	return Piece{}, false
}

func (board Board) String() string {
	var result string
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			square := Square{Rank: rank, File: file}
			piece, isOccupied := board.Get(square)

			if !isOccupied {
				result += " "
				continue
			}

			result += piece.Type.String()
		}

		result += "\n"
	}
	return result
}

func (board Board) IsOccupied(square Square) bool {
	_, isOccupied := board.Get(square)
	return isOccupied
}

func (board Board) IsOccupiedByEnemy(square Square, side Side) bool {
	piece, isOccupied := board.Get(square)
	return isOccupied && piece.Side != side
}
