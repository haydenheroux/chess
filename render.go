package main

import (
	gloss "github.com/charmbracelet/lipgloss"
)

var (
	blackSquare = gloss.NewStyle().Background(gloss.Color("#3E73D6")).Width(9).Height(3).PaddingTop(1).PaddingLeft(4)
	whiteSquare = blackSquare.Copy().Background(gloss.Color("#1466FF"))

	blackPiece = gloss.NewStyle().Foreground(gloss.Color("#000000"))
	whitePiece = gloss.NewStyle().Foreground(gloss.Color("#FFFFFF"))
)

func (m model) renderSquare(square Square) string {
	var squareStyle gloss.Style

	if square.IsBlack() {
		squareStyle = blackSquare
	} else {
		squareStyle = whiteSquare
	}

	piece, isOccupied := m.board.Get(square)

	thisPiece := " "

	if isOccupied {
		pieceType := piece.Type.String()

		var pieceStyle gloss.Style

		if piece.Side == Black {
			pieceStyle = blackPiece
		} else {
			pieceStyle = whitePiece
		}

		thisPiece = pieceStyle.Render(pieceType)
	}

	return squareStyle.Render(thisPiece)
}

func (m model) renderBoard(side Side) string {
	var board [][]string

	for rank := 7; rank >= 0; rank-- {

		var thisRank []string

		for file := 0; file < 8; file++ {
			square := Square{Rank: rank, File: file}

			thisRank = append(thisRank, m.renderSquare(square))
		}

		board = append(board, thisRank)
	}

	if side == Black {
		flip(board)
	}

	var view string

	for _, rank := range board {
		view += gloss.JoinHorizontal(gloss.Top, rank...)
		view += "\n"
	}

	return view
}
