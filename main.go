package main

import (
	tea "github.com/charmbracelet/bubbletea"
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
