package main

import "errors"
import "fmt"

type Square struct {
	Rank int
	File int
}

func Notation(notation string) (Square, error) {
	if len(notation) != 2 {
		return Square{}, errors.New("Input does not represent a square on the board.")
	}

	rank := int(notation[1]-'0') - 1
	file := int(notation[0] - 'a')

	if rank < 0 || rank > 7 || file < 0 || file > 7 {
		return Square{}, errors.New("Input does not represent a square on the board.")
	} 

	return Square{Rank: rank, File: file}, nil
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
	Type   PieceType
	Side   Side
}

func NewPiece(pieceType PieceType, side Side) Piece {
	return Piece{
		Type:   pieceType,
		Side:   side,
	}
}

func NewPawn(side Side) Piece {
	return NewPiece(Pawn, side)
}

func NewBishop(side Side) Piece {
	return NewPiece(Bishop, side)
}

func NewKnight(side Side) Piece {
	return NewPiece(Knight, side)
}

func NewRook(side Side) Piece {
	return NewPiece(Rook, side)
}

func NewQueen(side Side) Piece {
	return NewPiece(Queen, side)
}

func NewKing(side Side) Piece {
	return NewPiece(King, side)
}
