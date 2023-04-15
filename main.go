package main

import (
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	board Board
	side Side
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

func (m model) renderWhiteView() string {
	var result string
	for rank := 7; rank >= 0; rank-- {

		var thisRank []string

		for file := 0; file < 8; file++ {
			square := Square{Rank: rank, File: file}

			thisRank = append(thisRank, m.renderSquare(square))
		}

		result += gloss.JoinHorizontal(gloss.Top, thisRank...)
		result += "\n"
	}
	
	return result
}

func (m model) renderBlackView() string {
	var result string
	for rank := 0; rank < 8; rank++ {

		var thisRank []string

		for file := 7; file >= 0; file-- {
			square := Square{Rank: rank, File: file}

			thisRank = append(thisRank, m.renderSquare(square))
		}

		result += gloss.JoinHorizontal(gloss.Top, thisRank...)
		result += "\n"
	}
	
	return result
}

func (m model) View() string {
	if m.side == White {
		return m.renderWhiteView()
	} else {
		return m.renderBlackView()
	}
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
