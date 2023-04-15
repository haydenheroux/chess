package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	board Board
	side  Side
	textInput textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			m.textInput.SetValue("")
			if m.side == White {
				m.side = Black
			} else {
				m.side = White
			}

		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.renderBoard(m.side) + m.textInput.View()
}

func initialModel() model {

	ti := textinput.New()
	ti.Placeholder = "your move..."
	ti.Focus()

	return model{
		board: NewBoard(), 
		side: White,
		textInput: ti,
	}
}

func main() {
	m := initialModel()

	m.board = append(m.board, NewPawn(Notation("d6"), Black))

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		//fmt.Printf("Alas, there's been an error: %v", err)
		//os.Exit(1)
	}
}
