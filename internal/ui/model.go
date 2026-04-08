package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vaisakh678/termtype/internal/history"
	"github.com/vaisakh678/termtype/internal/words"
)

const defaultWordCount = 30

var (
	correctStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6e3a1"))
	incorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8"))
	cursorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#cdd6f4")).Underline(true)
	dimStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#6c7086"))
	titleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#89b4fa")).Bold(true)
	statStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#f9e2af")).Bold(true)
	labelStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6adc8"))
	selectedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#cba6f7")).Bold(true)
)

type state int

const (
	stateMenu state = iota
	stateReady
	stateTyping
	stateDone
)

type mode struct {
	name     string
	label    string
	duration time.Duration // 0 = word mode (no timer)
}

var modes = []mode{
	{name: "words", label: "30 words", duration: 0},
	{name: "15s", label: "15 seconds", duration: 15 * time.Second},
	{name: "30s", label: "30 seconds", duration: 30 * time.Second},
	{name: "60s", label: "60 seconds", duration: 60 * time.Second},
}

type tickMsg time.Time

type Model struct {
	words       []string
	input       []rune
	state       state
	start       time.Time
	elapsed     time.Duration
	wpm         float64
	liveWPM     float64
	accuracy    float64
	width       int
	modeIndex   int
	timeLimit   time.Duration
	history     history.History
}

func New() Model {
	return Model{
		state:   stateMenu,
		history: history.Load(),
	}
}

func (m *Model) startTest() {
	md := modes[m.modeIndex]
	m.timeLimit = md.duration
	if md.duration > 0 {
		// Timer mode: generate lots of words
		m.words = words.Generate(200)
	} else {
		m.words = words.Generate(defaultWordCount)
	}
	m.input = nil
	m.state = stateReady
	m.wpm = 0
	m.liveWPM = 0
	m.accuracy = 0
	m.elapsed = 0
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
			m.elapsed = time.Since(m.start)
			m.updateLiveWPM()

			// Timer mode: check if time is up
			if m.timeLimit > 0 && m.elapsed >= m.timeLimit {
				m.elapsed = m.timeLimit
				m.state = stateDone
				m.calculateStats()
				m.saveResult()
				return m, nil
			}
			return m, tick()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			if m.state == stateMenu {
				return m, tea.Quit
			}
			// Go back to menu from other states
			m.state = stateMenu
			return m, nil

		case "ctrl+r":
			if m.state != stateMenu {
				m.startTest()
				return m, nil
			}
		}

		// Menu navigation
		if m.state == stateMenu {
			switch msg.String() {
			case "up", "k":
				if m.modeIndex > 0 {
					m.modeIndex--
				}
			case "down", "j":
				if m.modeIndex < len(modes)-1 {
					m.modeIndex++
				}
			case "enter", " ":
				m.startTest()
			case "q":
				return m, tea.Quit
			}
			return m, nil
		}

		// Typing input
		switch msg.String() {
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil

		case "ctrl+w", "alt+backspace":
			// Delete whole word
			if len(m.input) > 0 {
				// Skip trailing spaces
				i := len(m.input) - 1
				for i >= 0 && m.input[i] == ' ' {
					i--
				}
				// Delete back to previous space
				for i >= 0 && m.input[i] != ' ' {
					i--
				}
				m.input = m.input[:i+1]
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

			// Check if test is complete (word mode only)
			if m.timeLimit == 0 {
				target := strings.Join(m.words, " ")
				if len(m.input) >= len(target) {
					m.state = stateDone
					m.elapsed = time.Since(m.start)
					m.calculateStats()
					m.saveResult()
					return m, nil
				}
			}
			return m, tick()
		}
	}

	return m, nil
}

func (m *Model) updateLiveWPM() {
	target := []rune(strings.Join(m.words, " "))
	correct := 0
	for i, ch := range m.input {
		if i < len(target) && ch == target[i] {
			correct++
		}
	}
	minutes := m.elapsed.Minutes()
	if minutes > 0 {
		m.liveWPM = (float64(correct) / 5.0) / minutes
	}
}

func (m *Model) calculateStats() {
	target := []rune(strings.Join(m.words, " "))
	correct := 0
	for i, ch := range m.input {
		if i < len(target) && ch == target[i] {
			correct++
		}
	}

	minutes := m.elapsed.Minutes()
	if minutes == 0 {
		return
	}

	m.wpm = (float64(correct) / 5.0) / minutes
	if len(m.input) > 0 {
		m.accuracy = (float64(correct) / float64(len(m.input))) * 100
	}
}

func (m *Model) saveResult() {
	md := modes[m.modeIndex]
	m.history.Add(history.Result{
		WPM:      m.wpm,
		Accuracy: m.accuracy,
		Duration: m.elapsed.Seconds(),
		Mode:     md.name,
		Date:     time.Now(),
	})
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
