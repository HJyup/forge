package tui

import (
	"github.com/HJyup/forge/config"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	BluePrimary   = lipgloss.AdaptiveColor{Light: "#007BFF", Dark: "#1B84F3"}
	BlueAccent    = lipgloss.AdaptiveColor{Light: "#4DAAFD", Dark: "#4FB0FF"}
	BlueHighlight = lipgloss.AdaptiveColor{Light: "#339CFF", Dark: "#61BFFF"}
	BlueBorder    = lipgloss.AdaptiveColor{Light: "#005FCC", Dark: "#3A91F5"}

	TextPrimary    = lipgloss.AdaptiveColor{Light: "#121212", Dark: "#F2F2F2"}
	TextSecondary  = lipgloss.AdaptiveColor{Light: "#5A5A5F", Dark: "#A8A8AD"}
	TextDimmed     = lipgloss.AdaptiveColor{Light: "#9D9DA2", Dark: "#686868"}
	TextDimmedDesc = lipgloss.AdaptiveColor{Light: "#C5C5C9", Dark: "#505055"}

	GrayMuted       = lipgloss.AdaptiveColor{Light: "#D0D0D4", Dark: "#5C5C5E"}
	GrayHelp        = lipgloss.AdaptiveColor{Light: "#B4B4BA", Dark: "#8A8A90"}
	GraySubdued     = lipgloss.AdaptiveColor{Light: "#A0A0A5", Dark: "#737377"}
	GrayVerySubdued = lipgloss.AdaptiveColor{Light: "#EAEAED", Dark: "#3F3F42"}
)

func NewRootTUIStyles() (s list.Styles) {
	s.TitleBar = lipgloss.NewStyle().Padding(config.TitleBarPaddingV, config.TitleBarPaddingH, config.TitleBarPaddingV, config.ItemPaddingLeft)

	s.Title = lipgloss.NewStyle().
		Bold(true).
		Background(BluePrimary).
		Padding(config.TitlePadding, config.TitlePadding)

	s.Spinner = lipgloss.NewStyle().Foreground(TextSecondary)
	s.FilterPrompt = lipgloss.NewStyle().Foreground(BlueAccent)
	s.FilterCursor = lipgloss.NewStyle().Foreground(BlueBorder)

	s.DefaultFilterCharacterMatch = lipgloss.NewStyle().
		Underline(true).
		Foreground(BlueHighlight)

	s.StatusBar = lipgloss.NewStyle().
		Foreground(TextSecondary).
		Padding(config.StatusPaddingV, config.StatusPaddingH, config.TitleBarPaddingV, config.ItemPaddingLeft)

	s.StatusEmpty = lipgloss.NewStyle().Foreground(GraySubdued)
	s.StatusBarActiveFilter = lipgloss.NewStyle().Foreground(BlueHighlight)
	s.StatusBarFilterCount = lipgloss.NewStyle().Foreground(GrayVerySubdued)

	s.NoItems = lipgloss.NewStyle().Foreground(GrayMuted)
	s.PaginationStyle = lipgloss.NewStyle().PaddingLeft(config.ItemPaddingLeft)

	s.HelpStyle = lipgloss.NewStyle().
		Padding(config.HelpPaddingTop, config.StatusPaddingH, config.StatusPaddingV, config.HelpPaddingLeft).
		Foreground(GrayHelp)

	s.ActivePaginationDot = lipgloss.NewStyle().
		Foreground(BlueHighlight).
		SetString("•")

	s.InactivePaginationDot = lipgloss.NewStyle().
		Foreground(GrayVerySubdued).
		SetString("•")

	s.DividerDot = lipgloss.NewStyle().
		Foreground(GrayVerySubdued).
		SetString(" • ")

	return s
}

func NewListItemStyles() (s list.DefaultItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(TextPrimary).
		Padding(config.StatusPaddingV, config.StatusPaddingH, config.StatusPaddingV, config.ItemPaddingLeft)

	s.NormalDesc = s.NormalTitle.Foreground(TextSecondary)

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(BluePrimary).
		Foreground(BlueBorder).
		Bold(true).
		Padding(config.StatusPaddingV, config.StatusPaddingH, config.StatusPaddingV, config.BorderPaddingLeft)

	s.SelectedDesc = s.SelectedTitle.Foreground(BlueAccent)

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(TextDimmed).
		Padding(config.StatusPaddingV, config.StatusPaddingH, config.StatusPaddingV, config.ItemPaddingLeft)

	s.DimmedDesc = s.DimmedTitle.Foreground(TextDimmedDesc)

	s.FilterMatch = lipgloss.NewStyle().
		Underline(true).
		Foreground(BlueHighlight)

	return s
}
