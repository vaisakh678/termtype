package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString(titleStyle.Render("  termtype"))
	b.WriteString("\n\n")

	switch m.state {
	case stateReady:
		b.WriteString(m.renderText())
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("  start typing to begin..."))

	case stateTyping:
		b.WriteString(m.renderText())
		b.WriteString("\n\n")
		elapsed := m.duration.Seconds()
		b.WriteString(dimStyle.Render(fmt.Sprintf("  %.1fs", elapsed)))

	case stateDone:
		b.WriteString(m.renderText())
		b.WriteString("\n\n")
		b.WriteString(statStyle.Render(fmt.Sprintf("  %.0f wpm", m.wpm)))
		b.WriteString(labelStyle.Render("  |  "))
		b.WriteString(statStyle.Render(fmt.Sprintf("%.1f%% accuracy", m.accuracy)))
		b.WriteString(labelStyle.Render("  |  "))
		b.WriteString(statStyle.Render(fmt.Sprintf("%.1fs", m.duration.Seconds())))
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("  ctrl+r to restart  |  esc to quit"))
	}

	b.WriteString("\n")
	return b.String()
}

func (m Model) renderText() string {
	target := []rune(strings.Join(m.words, " "))
	var b strings.Builder

	b.WriteString("  ")

	for i, ch := range target {
		char := string(ch)
		if i < len(m.input) {
			if m.input[i] == ch {
				b.WriteString(correctStyle.Render(char))
			} else {
				b.WriteString(incorrectStyle.Render(char))
			}
		} else if i == len(m.input) {
			b.WriteString(cursorStyle.Render(char))
		} else {
			b.WriteString(dimStyle.Render(char))
		}
	}

	return b.String()
}
