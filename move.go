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
		moves := board.ranksUntilLimit(square, direction, limit)
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
		return board.legalBishopMoves(square, piece.Side)
	case Knight:
		return []Square{}
	case Rook:
		return board.legalRookMoves(square, piece.Side)
	case Queen:
		bishop := board.legalBishopMoves(square, piece.Side)
		rook := board.legalRookMoves(square, piece.Side)
		return append(bishop, rook...)
	case King:
		potentialMoves := []Square{
			{Rank: square.Rank + int(Positive)},
			{Rank: square.Rank + int(Negative)},
			{File: square.File + int(Positive)},
			{File: square.File + int(Negative)},
			{Rank: square.Rank + int(Positive), File: square.File + int(Positive)},
			{Rank: square.Rank + int(Negative), File: square.File + int(Positive)},
			{Rank: square.Rank + int(Negative), File: square.File + int(Negative)},
			{Rank: square.Rank + int(Positive), File: square.File + int(Negative)},
		}
		moves := []Square{}

		for _, move := range potentialMoves {
			if !board.IsAlly(move, piece.Side) {
				moves = append(moves, move)
			}
		}

		return moves
	}

	return []Square{}
}

func (board Board) legalBishopMoves(from Square, side Side) []Square {
	pp := board.diagonalUntilEnemy(from, Positive, Positive, side)
	np := board.diagonalUntilEnemy(from, Negative, Positive, side)
	nn := board.diagonalUntilEnemy(from, Negative, Negative, side)
	pn := board.diagonalUntilEnemy(from, Positive, Negative, side)
	moves := append(pp, np...)
	moves = append(moves, nn...)
	moves = append(moves, pn...)
	return moves
}

func (board Board) legalRookMoves(from Square, side Side) []Square {
	pr := board.ranksUntilEnemy(from, Positive, side)
	nr := board.ranksUntilEnemy(from, Negative, side)
	pf := board.filesUntilEnemy(from, Positive, side)
	nf := board.filesUntilEnemy(from, Negative, side)
	moves := append(pr, nr...)
	moves = append(moves, pf...)
	moves = append(moves, nf...)
	return moves
}

func (board Board) ranksUntilEnemy(from Square, direction Direction, side Side) []Square {
	var squares []Square
	init := from.Rank + int(direction)
	for rank := init; rank >= 0 && rank < 8; rank += int(direction) {
		square := Square{Rank: rank, File: from.File}
		if board.IsAlly(square, side) {
			break
		}
		squares = append(squares, square)
		if board.IsEnemy(square, side) {
			break
		}
	}
	return squares
}

func (board Board) filesUntilEnemy(from Square, direction Direction, side Side) []Square {
	var squares []Square
	init := from.File + int(direction)
	for file := init; file >= 0 && file < 8; file += int(direction) {
		square := Square{Rank: from.Rank, File: file}
		if board.IsAlly(square, side) {
			break
		}
		squares = append(squares, square)
		if board.IsEnemy(square, side) {
			break
		}
	}
	return squares
}

func (board Board) diagonalUntilEnemy(from Square, rankDirection Direction, fileDirection Direction, side Side) []Square {
	var squares []Square
	initRank := from.Rank + int(rankDirection)
	initFile := from.File + int(fileDirection)

	for rank, file := initRank, initFile; rank >= 0 && rank < 8 && file >= 0 && file < 8; rank, file = rank + int(rankDirection), file + int(fileDirection) {
		square := Square{Rank: rank, File: file}
		if board.IsAlly(square, side) {
			return squares
		}
		squares = append(squares, square)
		if board.IsEnemy(square, side) {
			return squares
		}
	}
	return squares
}

func (board Board) ranksUntilLimit(from Square, direction Direction, limit int) []Square {
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

func (board Board) filesUntilLimit(from Square, direction Direction, limit int) []Square {
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
