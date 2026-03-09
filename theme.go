package diag

import "charm.land/lipgloss/v2"

type Theme struct {
	Error   lipgloss.Style
	Warning lipgloss.Style
	Message lipgloss.Style
	Help    lipgloss.Style
	Note    lipgloss.Style
	Muted   lipgloss.Style
}

func DefaultTheme() Theme {
	return Theme{
		Error:   lipgloss.NewStyle().Foreground(lipgloss.Color("#EB4268")).Bold(true),
		Warning: lipgloss.NewStyle().Foreground(lipgloss.Color("#E8FE96")).Bold(true),
		Message: lipgloss.NewStyle().Bold(true),
		Help:    lipgloss.NewStyle().Foreground(lipgloss.Color("#00A4FF")),
		Note:    lipgloss.NewStyle().Foreground(lipgloss.Color("#858392")),
		Muted:   lipgloss.NewStyle().Foreground(lipgloss.Color("#858392")),
	}
}
