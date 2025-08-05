package terminal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type PrettyPrinter struct {
	green     lipgloss.Color
	red       lipgloss.Color
	blue      lipgloss.Color
	yellow    lipgloss.Color
	gray      lipgloss.Color
	lightGray lipgloss.Color

	successStyle lipgloss.Style
	errorStyle   lipgloss.Style
	infoStyle    lipgloss.Style
	printStyle   lipgloss.Style
	warningStyle lipgloss.Style
	diffAddStyle lipgloss.Style
	diffDelStyle lipgloss.Style
	headerStyle  lipgloss.Style
}

func NewPrettyPrinter() *PrettyPrinter {
	green := lipgloss.Color("42")
	red := lipgloss.Color("196")
	blue := lipgloss.Color("39")
	yellow := lipgloss.Color("220")
	gray := lipgloss.Color("245")
	lightGray := lipgloss.Color("241")

	return &PrettyPrinter{
		green:     green,
		red:       red,
		blue:      blue,
		yellow:    yellow,
		gray:      gray,
		lightGray: lightGray,

		successStyle: lipgloss.NewStyle().Foreground(green).Bold(true),
		errorStyle:   lipgloss.NewStyle().Foreground(red).Bold(true),
		infoStyle:    lipgloss.NewStyle().Foreground(blue).Bold(true),
		printStyle:   lipgloss.NewStyle().Foreground(gray),
		warningStyle: lipgloss.NewStyle().Foreground(yellow).Bold(true),
		diffAddStyle: lipgloss.NewStyle().Foreground(green),
		diffDelStyle: lipgloss.NewStyle().Foreground(red),
		headerStyle:  lipgloss.NewStyle().Foreground(blue).Bold(true).Underline(true),
	}
}

func (pp *PrettyPrinter) Success(message string, keyvals ...any) string {
	return pp.successStyle.Render("✅ " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Print(message string, keyvals ...any) string {
	return pp.printStyle.Render(pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Error(message string, keyvals ...any) string {
	return pp.errorStyle.Render("❌ " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Info(message string, keyvals ...any) string {
	return pp.infoStyle.Render("ℹ️ " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Warning(message string, keyvals ...any) string {
	return pp.warningStyle.Render("⚠️ " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Header(title string) string {
	return pp.headerStyle.Render(title)
}

func (pp *PrettyPrinter) Box(title, content string) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(pp.blue).
		Padding(1, 2)

	if title != "" {
		content = pp.headerStyle.Render(title) + "\n\n" + content
	}

	return boxStyle.Render(content)
}

func (pp *PrettyPrinter) List(items []string) string {
	var output strings.Builder

	for _, item := range items {
		output.WriteString(pp.infoStyle.Render("• "))
		output.WriteString(item)
		output.WriteString("\n")
	}

	return output.String()
}

func (pp *PrettyPrinter) KeyValue(pairs map[string]string) string {
	var output strings.Builder

	for key, value := range pairs {
		output.WriteString(pp.infoStyle.Render(key + ": "))
		output.WriteString(value)
		output.WriteString("\n")
	}

	return output.String()
}

func (pp *PrettyPrinter) Separator() string {
	return lipgloss.NewStyle().
		Foreground(pp.gray).
		Render(strings.Repeat("─", 50))
}

func (pp *PrettyPrinter) formatMessage(message string, keyvals ...any) string {
	if len(keyvals) == 0 {
		return message
	}

	var parts []string
	parts = append(parts, message)

	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			key := fmt.Sprintf("%v", keyvals[i])
			value := fmt.Sprintf("%v", keyvals[i+1])
			parts = append(parts, fmt.Sprintf("%s=%s", key, value))
		} else {
			key := fmt.Sprintf("%v", keyvals[i])
			parts = append(parts, fmt.Sprintf("%s=<missing>", key))
		}
	}

	return strings.Join(parts, " ")
}
