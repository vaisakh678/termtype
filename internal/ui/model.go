package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vaisakh678/termtype/internal/words"
)

const wordCount = 30

var (
	correctStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6e3a1"))
	incorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8"))
	cursorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#cdd6f4")).Underline(true)
	dimStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#6c7086"))
	titleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#89b4fa")).Bold(true)
	statStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#f9e2af")).Bold(true)
	labelStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6adc8"))
)

type state int

const (
	stateReady state = iota
	stateTyping
	stateDone
)

type tickMsg time.Time

type Model struct {
	words    []string
	input    []rune
	state    state
	start    time.Time
	duration time.Duration
	wpm      float64
	accuracy float64
	width    int
}

func New() Model {
	return Model{
		words: words.Generate(wordCount),
		state: stateReady,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil

	case tickMsg:
		if m.state == stateTyping {
			m.duration = time.Since(m.start)
			return m, tick()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "ctrl+r":
			m = New()
			m.state = stateReady
			return m, nil

		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil
		}

		if m.state == stateDone {
			return m, nil
		}

		// Single character input
		if len(msg.String()) == 1 {
			if m.state == stateReady {
				m.state = stateTyping
				m.start = time.Now()
			}

			m.input = append(m.input, rune(msg.String()[0]))

			// Check if test is complete
			target := strings.Join(m.words, " ")
			if len(m.input) >= len(target) {
				m.state = stateDone
				m.duration = time.Since(m.start)
				m.calculateStats()
				return m, nil
			}
			return m, tick()
		}
	}

	return m, nil
}

func (m *Model) calculateStats() {
	target := []rune(strings.Join(m.words, " "))
	correct := 0
	for i, ch := range m.input {
		if i < len(target) && ch == target[i] {
			correct++
		}
	}

	minutes := m.duration.Minutes()
	if minutes == 0 {
		return
	}

	// WPM = (correct chars / 5) / minutes
	m.wpm = (float64(correct) / 5.0) / minutes
	m.accuracy = (float64(correct) / float64(len(m.input))) * 100
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
