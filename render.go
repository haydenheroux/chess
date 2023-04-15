package main

import (
	gloss "github.com/charmbracelet/lipgloss"
)

var (
	blackSquare = gloss.NewStyle().Width(9).Height(3).PaddingTop(1).PaddingLeft(4).Background(gloss.Color("#6272a4"))
	whiteSquare = blackSquare.Copy().Background(gloss.Color("#ff79c6"))

	blackPiece = gloss.NewStyle().Foreground(gloss.Color("#000000")).Bold(true)
	whitePiece = gloss.NewStyle().Foreground(gloss.Color("#ffffff")).Bold(true)

	inputBox = gloss.NewStyle().Background(gloss.Color("#6272a4")).Foreground(gloss.Color("#f8f8f2")).Width(70).Height(3).PaddingTop(1).PaddingLeft(1).BorderStyle(gloss.NormalBorder()).BorderForeground(gloss.Color("#f8f8f2")).BorderBackground(gloss.Color("#6272a4")).MarginTop(1)

	errorBox = inputBox.Copy().Background(gloss.Color("#ff5555")).BorderBackground(gloss.Color("#ff5555")).MarginTop(1).Bold(true)
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

func (m model) renderBoard() string {
	var board [][]string

	for rank := 7; rank >= 0; rank-- {

		var thisRank []string

		for file := 0; file < 8; file++ {
			square := Square{Rank: rank, File: file}

			thisRank = append(thisRank, m.renderSquare(square))
		}

		board = append(board, thisRank)
	}

	if m.side == Black {
		flip(board)
	}

	var view string

	for _, rank := range board {
		view += gloss.JoinHorizontal(gloss.Top, rank...)
		view += "\n"
	}

	return view
}

func (m model) renderInput() string {
	return inputBox.Render(m.textInput.View())
}

func (m model) renderError() string {
	return errorBox.Render(m.error.Error())
}
