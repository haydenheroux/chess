package main

import (
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	board Board
	side  Side
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			if m.side == White {
				m.side = Black
			} else {
				m.side = White
			}

		}
	}

	return m, nil
}

var blackSquare = gloss.NewStyle().Background(gloss.Color("#3E73D6")).Width(9).Height(3).PaddingTop(1).PaddingLeft(4)
var whiteSquare = blackSquare.Copy().Background(gloss.Color("#1466FF"))

var blackPiece = gloss.NewStyle().Foreground(gloss.Color("#000000"))
var whitePiece = gloss.NewStyle().Foreground(gloss.Color("#FFFFFF"))

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

func flip(s0 [][]string) {
	for l, r := 0, len(s0)-1; l < r; l, r = l+1, r-1 {
		s0[l], s0[r] = s0[r], s0[l]
	}

	for _, s1 := range s0 {
		reverse(s1)
	}
}

func reverse(s []string) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
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

func (m model) View() string {
	return m.renderBoard(m.side)
}

func main() {
	initialModel := model{board: NewBoard(), side: White}
	initialModel.board = append(initialModel.board, NewPawn(Notation("d6"), Black))

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		//fmt.Printf("Alas, there's been an error: %v", err)
		//os.Exit(1)
	}
}
