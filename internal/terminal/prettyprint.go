package terminal

import (
	"fmt"
	"strings"

	"gobsidian/internal/diff"

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
	return pp.infoStyle.Render(pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Error(message string, keyvals ...any) string {
	return pp.errorStyle.Render("❌  " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Info(message string, keyvals ...any) string {
	return pp.infoStyle.Render("ℹ️  " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Warning(message string, keyvals ...any) string {
	return pp.warningStyle.Render("⚠️  " + pp.formatMessage(message, keyvals...))
}

func (pp *PrettyPrinter) Header(title string) string {
	return pp.headerStyle.Render(title)
}

func (pp *PrettyPrinter) Diff(title, oldText, newText string) string {
	var output strings.Builder

	if title != "" {
		output.WriteString(pp.headerStyle.Render(title))
		output.WriteString("\n")
	}

	oldLines := strings.Split(oldText, "\n")
	newLines := strings.Split(newText, "\n")

	maxLines := max(len(oldLines), len(newLines))

	for i := 0; i < maxLines; i++ {
		var oldLine, newLine string

		if i < len(oldLines) {
			oldLine = oldLines[i]
		}
		if i < len(newLines) {
			newLine = newLines[i]
		}

		if oldLine != newLine {
			if oldLine != "" {
				output.WriteString(pp.diffDelStyle.Render("- " + oldLine))
				output.WriteString("\n")
			}
			if newLine != "" {
				output.WriteString(pp.diffAddStyle.Render("+ " + newLine))
				output.WriteString("\n")
			}
		} else if oldLine != "" {
			output.WriteString(lipgloss.NewStyle().Foreground(pp.gray).Render("  " + oldLine))
			output.WriteString("\n")
		}
	}

	return output.String()
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

func (pp *PrettyPrinter) PrintDiffs(diffs []diff.Diff) string {
	if len(diffs) == 0 {
		return pp.Info("No changes detected")
	}

	var output strings.Builder

	output.WriteString(pp.headerStyle.Render("📄 FILE CHANGES"))
	output.WriteString("\n\n")

	changesByType := make(map[diff.ChangeType][]diff.Diff)
	for _, d := range diffs {
		changesByType[d.ChangeType] = append(changesByType[d.ChangeType], d)
	}

	for changeType, changes := range changesByType {
		if len(changes) == 0 {
			continue
		}

		var icon, color string
		var style lipgloss.Style

		switch changeType {
		case diff.Added:
			icon = "+"
			color = "42" // green
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		case diff.Modified:
			icon = "~"
			color = "220" // yellow
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		case diff.Deleted:
			icon = "-"
			color = "196" // red
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		}

		sectionHeader := fmt.Sprintf("%s %s (%d)", icon, changeType.String(), len(changes))
		output.WriteString(style.Bold(true).Render(sectionHeader))
		output.WriteString("\n")

		for _, change := range changes {
			fileEntry := fmt.Sprintf("  %s %s", icon, change.File)
			output.WriteString(style.Render(fileEntry))
			output.WriteString("\n")
		}
		output.WriteString("\n")
	}

	summary := fmt.Sprintf("Total changes: %d", len(diffs))
	output.WriteString(pp.infoStyle.Render(summary))

	return output.String()
}

func (pp *PrettyPrinter) PrintDiffsSummary(diffs []diff.Diff) string {
	if len(diffs) == 0 {
		return pp.Info("No changes detected")
	}

	counts := make(map[diff.ChangeType]int)
	for _, d := range diffs {
		counts[d.ChangeType]++
	}

	var parts []string
	if counts[diff.Added] > 0 {
		parts = append(parts, lipgloss.NewStyle().
			Foreground(lipgloss.Color("42")).
			Render(fmt.Sprintf("+%d", counts[diff.Added])))
	}
	if counts[diff.Modified] > 0 {
		parts = append(parts, lipgloss.NewStyle().
			Foreground(lipgloss.Color("220")).
			Render(fmt.Sprintf("~%d", counts[diff.Modified])))
	}
	if counts[diff.Deleted] > 0 {
		parts = append(parts, lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Render(fmt.Sprintf("-%d", counts[diff.Deleted])))
	}

	summary := fmt.Sprintf("📄 Changes: %s", strings.Join(parts, " "))
	return pp.infoStyle.Render(summary)
}
