package main

import (
	"errors"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

type state int

const (
	wantSelection state = iota
	wantMove
)

func (s state) String() string {
	return []string{"Select a piece to move", "Select the position to move to"}[s]
}

type model struct {
	board     Board
	side      Side
	textInput textinput.Model
	state     state
	selection Square
	error error
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
			switch m.state {

			case wantSelection:
				var selection Square
				selection, m.error = Notation(text)
				if m.error == nil {
					if m.board.IsAlly(selection, m.side) {
						m.selection = selection
						m.state = wantMove
						m.textInput.Prompt = "Select a square to move to: "
						m.error = nil
					} else {
						m.error = errors.New("Inputted square does not contain an ally.")
					}
				}

			case wantMove:
				var move Square
				move, m.error = Notation(text)
				if m.error == nil {
					piece, _ := m.board.Get(m.selection)
					if piece.IsLegalMove(move) {
						m.error = m.board.Move(m.selection, move)
						if m.error == nil {
							m.state = wantSelection
							m.textInput.Prompt = "Select a piece to move: "
						}
					} else {
						m.error = errors.New("Inputted move was not legal.")
					}
				}
			}

		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	view := m.renderBoard() + m.renderInput()
	if m.error != nil {
		view += m.renderError()
	}
	return view
}

func initialModel() model {

	ti := textinput.New()
	ti.Prompt = "Select a piece to move: "
	ti.PromptStyle = gloss.NewStyle().Foreground(gloss.Color("#f8f8f2")).Background(gloss.Color("#6272a4")).Bold(true)
	ti.PlaceholderStyle = ti.PromptStyle.Copy().Foreground(gloss.Color("#f8f8f2")).Background(gloss.Color("#6272a4")).Bold(false)
	ti.TextStyle = ti.PromptStyle.Copy().Foreground(gloss.Color("#f8f8f2")).Background(gloss.Color("#6272a4")).Bold(false)

	ti.Focus()

	return model{
		board:     NewBoard(),
		side:      White,
		textInput: ti,
		state:     wantSelection,
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
