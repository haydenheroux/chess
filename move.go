package main

import "math"

type Direction int

const (
	Positive Direction = 1
	Negative = -1 
)

func (board Board) IsLegalMove(piece Piece, from Square, to Square) bool {
	for _, move := range board.LegalMoves(piece, from) {
		if move == to {
			return true
		}
	}

	return false
}

func (board Board) LegalMoves(piece Piece, square Square) []Square {
	switch (piece.Type) {
	case Pawn:
		var direction Direction
		if piece.Side == Black {
			direction = Negative
		} else {
			direction = Positive
		}

		var limit int
		if piece.HasMoved {
			limit = 1
		} else {
			limit = 2
		}
		moves := board.verticalSquaresUntilObstructedWithLimit(square, direction, limit)
		diagonalPositive := Square{Rank: square.Rank + int(direction), File: square.File + int(Positive)}
		diagonalNegative := Square{Rank: square.Rank + int(direction), File: square.File + int(Negative)}
		if board.IsEnemy(diagonalPositive, piece.Side) {
			moves = append(moves, diagonalPositive)
		}
		if board.IsEnemy(diagonalNegative, piece.Side) {
			moves = append(moves, diagonalNegative)
		}
		return moves
	case Bishop:
	case Knight:
	case Rook:
	case Queen:
	case King:
		return []Square{}
	}

	return []Square{}
}

func (board Board) verticalSquaresUntilObstructed(from Square, direction Direction) []Square {
	var squares []Square
	init := from.Rank + int(direction)
	for rank := init; rank >= 0 && rank < 8; rank += int(direction) {
		square := Square{Rank: rank, File: from.File}
		if board.IsNotEmpty(square) {
			break
		}
		squares = append(squares, square)
	}
	return squares
}

func (board Board) horizontalSquaresUntilObstructed(from Square, direction Direction) []Square {
	var squares []Square
	init := from.File + int(direction)
	for file := init; file >= 0 && file < 8; file += int(direction) {
		square := Square{Rank: from.Rank, File: file}
		if board.IsNotEmpty(square) {
			break
		}
		squares = append(squares, square)
	}
	return squares
}

func (board Board) verticalSquaresUntilObstructedWithLimit(from Square, direction Direction, limit int) []Square {
	var squares []Square
	init := from.Rank + int(direction)
	for rank := init; rank >= 0 && rank < 8; rank += int(direction) {
		square := Square{Rank: rank, File: from.File}
		distance := int(math.Abs(float64(rank - init)))
		if board.IsNotEmpty(square) || distance >= limit {
			break
		}
		squares = append(squares, square)
	}
	return squares
}

func (board Board) horizontalSquaresUntilObstructedWithLimit(from Square, direction Direction, limit int) []Square {
	var squares []Square
	init := from.File + int(direction)
	for file := init; file >= 0 && file < 8; file += int(direction) {
		square := Square{Rank: from.Rank, File: file}
		distance := int(math.Abs(float64(file - init)))
		if board.IsNotEmpty(square) || distance >= limit {
			break
		}
		squares = append(squares, square)
	}
	return squares
}
