package terminal

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GenerationStep struct {
	Name     string
	duration time.Duration
	count    int
	done     bool
}

type ProgressManager struct {
	steps    []GenerationStep
	current  int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
	err      error
	program  *tea.Program
}

type StepCompleteMsg struct {
	Step     int
	Duration time.Duration
	Count    int
	Err      error
}

var (
	stepStyle        = lipgloss.NewStyle()
	currentStepStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle        = lipgloss.NewStyle().Margin(1, 0)
	checkMark        = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	errorMark        = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).SetString("✗")
)

type StepDefinition struct {
	Name  string
	Emoji string
}

func CreateStep(name, emoji string) StepDefinition {
	return StepDefinition{
		Name:  name,
		Emoji: emoji,
	}
}

func NewProgressManager(stepDefs []StepDefinition) *ProgressManager {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	steps := make([]GenerationStep, len(stepDefs))
	for i, def := range stepDefs {
		stepName := def.Name
		if def.Emoji != "" {
			stepName = def.Name + " " + def.Emoji
		}
		steps[i] = GenerationStep{
			Name:     stepName,
			duration: 0,
			count:    0,
			done:     false,
		}
	}

	return &ProgressManager{
		steps:    steps,
		current:  0,
		spinner:  s,
		progress: p,
		done:     false,
	}
}

func (m *ProgressManager) Start() error {
	m.program = tea.NewProgram(m)
	_, err := m.program.Run()
	return err
}

func (m *ProgressManager) StartAsync() {
	m.program = tea.NewProgram(m)
	go func() {
		if _, err := m.program.Run(); err != nil {
			// currently just don't do anything
		}
	}()
}

func (m *ProgressManager) SendStepComplete(step int, duration time.Duration, count int, err error) {
	if m.program != nil {
		m.program.Send(StepCompleteMsg{
			Step:     step,
			Duration: duration,
			Count:    count,
			Err:      err,
		})
	}
}

func (m *ProgressManager) Quit() {
	if m.program != nil {
		m.program.Quit()
	}
}

func (m *ProgressManager) GetStepCount() int {
	return len(m.steps)
}

func (m *ProgressManager) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *ProgressManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case StepCompleteMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, tea.Quit
		}

		m.steps[msg.Step].done = true
		m.steps[msg.Step].duration = msg.Duration
		m.steps[msg.Step].count = msg.Count

		if m.current >= len(m.steps)-1 {
			m.done = true
			return m, tea.Quit
		}

		m.current++
		progressCmd := m.progress.SetPercent(float64(m.current) / float64(len(m.steps)))
		stepLine := stepStyle.Render(
			fmt.Sprintf("%s %s (%d items, %v)",
				checkMark,
				m.steps[msg.Step].Name,
				msg.Count,
				msg.Duration,
			))

		return m, tea.Batch(
			progressCmd,
			tea.Printf(stepLine),
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m *ProgressManager) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if m.done {
		return doneStyle.Render("✅ Site generation complete!\n")
	}

	n := len(m.steps)
	w := lipgloss.Width(fmt.Sprintf("%d", n))
	stepCount := fmt.Sprintf(" %*d/%*d", w, m.current+1, w, n)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+stepCount))

	stepName := currentStepStyle.Render(m.steps[m.current].Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render(stepName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+stepCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + stepCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
