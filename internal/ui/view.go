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
	case stateMenu:
		b.WriteString(m.renderMenu())

	case stateReady:
		b.WriteString(m.renderText())
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("  start typing to begin..."))

	case stateTyping:
		b.WriteString(m.renderText())
		b.WriteString("\n\n")
		b.WriteString(m.renderLiveStats())

	case stateDone:
		b.WriteString(m.renderText())
		b.WriteString("\n\n")
		b.WriteString(m.renderFinalStats())
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("  ctrl+r restart  |  esc menu"))
	}

	b.WriteString("\n")
	return b.String()
}

func (m Model) renderMenu() string {
	var b strings.Builder

	b.WriteString(labelStyle.Render("  select mode:"))
	b.WriteString("\n\n")

	for i, md := range modes {
		if i == m.modeIndex {
			b.WriteString(selectedStyle.Render(fmt.Sprintf("  > %s", md.label)))
		} else {
			b.WriteString(dimStyle.Render(fmt.Sprintf("    %s", md.label)))
		}
		// Show personal best next to each mode
		best := m.history.BestWPMByMode(md.name)
		if best > 0 {
			b.WriteString(dimStyle.Render(fmt.Sprintf("  (best: %.0f wpm)", best)))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(dimStyle.Render("  enter to start  |  q to quit"))

	return b.String()
}

func (m Model) renderLiveStats() string {
	var b strings.Builder
	b.WriteString("  ")

	if m.timeLimit > 0 {
		remaining := m.timeLimit - m.elapsed
		if remaining < 0 {
			remaining = 0
		}
		b.WriteString(statStyle.Render(fmt.Sprintf("%.0fs", remaining.Seconds())))
		b.WriteString(dimStyle.Render("  |  "))
	} else {
		b.WriteString(dimStyle.Render(fmt.Sprintf("%.1fs", m.elapsed.Seconds())))
		b.WriteString(dimStyle.Render("  |  "))
	}

	b.WriteString(statStyle.Render(fmt.Sprintf("%.0f wpm", m.liveWPM)))

	return b.String()
}

func (m Model) renderFinalStats() string {
	var b strings.Builder

	b.WriteString(statStyle.Render(fmt.Sprintf("  %.0f wpm", m.wpm)))
	b.WriteString(labelStyle.Render("  |  "))
	b.WriteString(statStyle.Render(fmt.Sprintf("%.1f%% accuracy", m.accuracy)))
	b.WriteString(labelStyle.Render("  |  "))
	b.WriteString(statStyle.Render(fmt.Sprintf("%.1fs", m.elapsed.Seconds())))

	// Show if it's a personal best
	md := modes[m.modeIndex]
	best := m.history.BestWPMByMode(md.name)
	if m.wpm >= best && m.wpm > 0 {
		b.WriteString(labelStyle.Render("  |  "))
		b.WriteString(selectedStyle.Render("new best!"))
	}

	return b.String()
}

func (m Model) renderText() string {
	target := []rune(strings.Join(m.words, " "))
	maxWidth := m.width - 4 // 2 char padding each side
	if maxWidth <= 0 {
		maxWidth = 76
	}

	var lines []string
	var line strings.Builder
	lineLen := 0

	for i, ch := range target {
		// Wrap at word boundaries
		if ch == ' ' && lineLen >= maxWidth {
			lines = append(lines, line.String())
			line.Reset()
			lineLen = 0
			continue
		}

		char := string(ch)
		var styled string
		if i < len(m.input) {
			if m.input[i] == ch {
				styled = correctStyle.Render(char)
			} else {
				styled = incorrectStyle.Render(char)
			}
		} else if i == len(m.input) {
			styled = cursorStyle.Render(char)
		} else {
			styled = dimStyle.Render(char)
		}

		line.WriteString(styled)
		lineLen++
	}

	if line.Len() > 0 {
		lines = append(lines, line.String())
	}

	var result strings.Builder
	for i, l := range lines {
		result.WriteString("  ")
		result.WriteString(l)
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
