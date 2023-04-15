package main

import (
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	board Board
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

		}
	}

	return m, nil
}

var blackSquare = gloss.NewStyle().Background(gloss.Color("#3E73D6")).Width(9).Height(3).PaddingTop(1).PaddingLeft(4)
var whiteSquare = blackSquare.Copy().Background(gloss.Color("#1466FF"))

var blackPiece = gloss.NewStyle().Foreground(gloss.Color("#000000"))
var whitePiece = gloss.NewStyle().Foreground(gloss.Color("#FFFFFF"))

func (m model) View() string {
	var result string
	for rank := 7; rank >= 0; rank-- {

		var thisRank []string

		for file := 0; file < 8; file++ {
			square := Square{Rank: rank, File: file}

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

			thisRank = append(thisRank, squareStyle.Render(thisPiece))
		}

		result += gloss.JoinHorizontal(gloss.Top, thisRank...)
		result += "\n"
	}
	
	return result
}

func main() {

	initialModel := model{board: NewBoard()}

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		//fmt.Printf("Alas, there's been an error: %v", err)
		//os.Exit(1)
	}
}
