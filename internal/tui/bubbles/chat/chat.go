package chat

import (
	"fmt"

	"github.com/Drax-1/tandem/internal/config"
	"github.com/Drax-1/tandem/internal/message"
	"github.com/Drax-1/tandem/internal/session"
	"github.com/Drax-1/tandem/internal/tui/styles"
	"github.com/Drax-1/tandem/internal/tui/theme"
	"github.com/Drax-1/tandem/internal/version"
	"github.com/charmbracelet/lipgloss"
)

type SendMsg struct {
	Text        string
	Attachments []message.Attachment
}

type SessionSelectedMsg = session.Session

type SessionClearedMsg struct{}

type EditorFocusMsg bool

func header(width int) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		logo(width),
		repo(width),
		" ",
		cwd(width),
	)
}

func logo(width int) string {
	logo := "Tandem"
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	versionText := baseStyle.
		Foreground(t.TextMuted()).
		Render(version.Version)

	return baseStyle.
		Bold(true).
		Width(width).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				logo,
				" ",
				versionText,
			),
		)
}

func repo(width int) string {
	repo := "https://github.com/Drax-1/tandem"
	t := theme.CurrentTheme()

	return styles.BaseStyle().
		Foreground(t.TextMuted()).
		Width(width).
		Render(repo)
}

func cwd(width int) string {
	cwd := fmt.Sprintf("cwd: %s", config.WorkingDirectory())
	t := theme.CurrentTheme()

	return styles.BaseStyle().
		Foreground(t.TextMuted()).
		Width(width).
		Render(cwd)
}
