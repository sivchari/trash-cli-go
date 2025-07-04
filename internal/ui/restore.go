package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sivchari/trash-cli-go/internal/trash"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	itemStyle = lipgloss.NewStyle().PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("#7D56F4")).
				Bold(true)

	paginationStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("#626262"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4)
)

type model struct {
	items       []trash.TrashItem
	cursor      int
	selected    map[int]struct{}
	quitting    bool
	err         error
	restored    []string
	multiSelect bool
	showHelp    bool
}

type tickMsg struct{}

func initialModel(items []trash.TrashItem) model {
	return model{
		items:       items,
		selected:    make(map[int]struct{}),
		multiSelect: false,
		showHelp:    true,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter", " ":
			if !m.multiSelect {
				item := m.items[m.cursor]
				if err := trash.RestoreFile(item); err != nil {
					m.err = err
				} else {
					m.restored = append(m.restored, item.OriginalPath)
					m.quitting = true
					return m, tea.Quit
				}
			} else {
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}

		case "tab":
			m.multiSelect = !m.multiSelect
			if !m.multiSelect {
				m.selected = make(map[int]struct{})
			}

		case "r":
			if m.multiSelect && len(m.selected) > 0 {
				for i := range m.selected {
					item := m.items[i]
					if err := trash.RestoreFile(item); err != nil {
						m.err = err
						break
					}
					m.restored = append(m.restored, item.OriginalPath)
				}
				m.quitting = true
				return m, tea.Quit
			}

		case "?":
			m.showHelp = !m.showHelp
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		if m.err != nil {
			return quitTextStyle.Render(fmt.Sprintf("Error: %v", m.err))
		}
		if len(m.restored) > 0 {
			restoredList := strings.Join(m.restored, "\n  ")
			return quitTextStyle.Render(fmt.Sprintf("✅ Restored:\n  %s", restoredList))
		}
		return quitTextStyle.Render("👋 No files restored.")
	}

	if len(m.items) == 0 {
		return quitTextStyle.Render("🗑️  Trash is empty")
	}

	var s strings.Builder

	title := "🗑️  Select files to restore"
	if m.multiSelect {
		title += fmt.Sprintf(" (Multi-select: %d selected)", len(m.selected))
	}
	s.WriteString(titleStyle.Render(title))
	s.WriteString("\n\n")

	for i, item := range m.items {
		cursor := "  "
		if m.cursor == i {
			cursor = "▶ "
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "✓"
		}

		display := fmt.Sprintf("%s %s%s",
			item.DeletionDate.Format("2006-01-02 15:04"),
			item.OriginalPath,
			"")

		if m.cursor == i {
			if m.multiSelect {
				s.WriteString(selectedItemStyle.Render(fmt.Sprintf("%s[%s] %s", cursor, checked, display)))
			} else {
				s.WriteString(selectedItemStyle.Render(fmt.Sprintf("%s%s", cursor, display)))
			}
		} else {
			if m.multiSelect {
				s.WriteString(itemStyle.Render(fmt.Sprintf("%s[%s] %s", cursor, checked, display)))
			} else {
				s.WriteString(itemStyle.Render(fmt.Sprintf("%s%s", cursor, display)))
			}
		}
		s.WriteString("\n")
	}

	if m.showHelp {
		s.WriteString("\n")
		help := ""
		if m.multiSelect {
			help = "Space: toggle • Tab: single mode • r: restore selected • ?: help • q: quit"
		} else {
			help = "Enter: restore • Tab: multi mode • ↑/↓: navigate • ?: help • q: quit"
		}
		s.WriteString(helpStyle.Render(help))
	}

	return s.String()
}

func RunRestoreUI(items []trash.TrashItem) error {
	p := tea.NewProgram(initialModel(items), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
