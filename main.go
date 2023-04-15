package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type state int

const (
	wantSelection state = iota
	wantMove 
)

func (s state) String() string {
	return []string{"wantSelection", "wantMove"}[s]
}

type model struct {
	board Board
	side  Side
	textInput textinput.Model
	state state
	selection Square
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
			text := m.textInput.Value()
			m.textInput.SetValue("")
			switch (m.state) {
			case wantSelection:
				selection := Notation(text)
				if m.board.IsOccupiedByAlly(selection, m.side) {
					m.selection = selection
					m.state = wantMove
				}
			case wantMove:
				move := Notation(text)
				isLegalMove := true
				if isLegalMove {
					_ = m.board.Move(m.selection, move)
					m.state = wantSelection
				}
			}

		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.renderBoard(m.side) + m.textInput.View() + "\n"
}

func initialModel() model {

	ti := textinput.New()
	ti.Placeholder = "your move..."
	ti.Focus()

	return model{
		board: NewBoard(), 
		side: White,
		textInput: ti,
		state: wantSelection,
		selection: Square{},
	}
}

func main() {
	m := initialModel()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		//fmt.Printf("Alas, there's been an error: %v", err)
		//os.Exit(1)
	}
}
