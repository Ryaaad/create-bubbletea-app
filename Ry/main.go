
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board    []string
	cursor   int
	Selected string
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
	
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
	
			case "down", "j":
				if m.cursor < len(m.board)-1 {
					m.cursor++
				}
	
			case "enter", " ":
				m.Selected = m.board[m.cursor]
				}
			}
			return m, nil
		}
		
func (m model) View() string {
	s := "\n"
			// Iterate over our choices
			for i := range m.board {
		
				cursor := " "
				if m.cursor == i {
					cursor = ">"
				}
				s += fmt.Sprintf("%s %s", cursor, m.board[i])
				s += "\n"
			}
			if m.Selected != "" {
				s += fmt.Sprintf("\nYou Selected %s \n", m.Selected)
			}
			// The footer
			s += "\nPress q to quit.\n"
		
			// Send the UI for rendering
			return s
		}
		
		func initialModel() model {
			return model{
				board:  []string{"A", "B", "C"},
				cursor: 0,
			}
		}
		
		func main() {
			p := tea.NewProgram(initialModel())
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
			fmt.Print("{.view}")
		}		
		
