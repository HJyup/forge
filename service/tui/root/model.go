package tui

import (
	"fmt"

	"github.com/HJyup/forge/config"
	"github.com/HJyup/forge/service/parser"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

var docStyle = lipgloss.NewStyle().Margin(config.DocMarginV, config.DocMarginH)

type ProjectItem struct {
	project parser.ProjectSummary
}

func (i ProjectItem) Title() string { return i.project.Name }
func (i ProjectItem) Description() string {
	timeStr := humanize.Time(i.project.LastModified)
	return fmt.Sprintf("%s%s%s", i.project.Path, config.ProjectSeparator, timeStr)
}
func (i ProjectItem) FilterValue() string { return i.project.Name }

type RootModel struct {
	list list.Model
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == config.KeyQuit {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m RootModel) View() string {
	return docStyle.Render(m.list.View())
}

func NewRootModel(projects []parser.ProjectSummary) RootModel {
	items := make([]list.Item, len(projects))
	for i, project := range projects {
		items[i] = ProjectItem{project: project}
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles = NewListItemStyles()

	l := list.New(items, delegate, 0, 0)
	l.Title = config.TUITitle
	l.Styles = NewRootTUIStyles()

	return RootModel{list: l}
}

func RunRootTUI(projects []parser.ProjectSummary) error {
	m := NewRootModel(projects)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}

	return nil
}
