package main

import "fmt"

type Square struct {
	Rank int
	File int
}

func Notation(notation string) Square {
	rank := int(notation[1]-'0') - 1
	file := int(notation[0] - 'a')
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
