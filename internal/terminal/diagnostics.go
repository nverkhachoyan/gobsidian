package terminal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type DiagnosticsSummary struct {
	TotalNotes            int
	NotesWithWarnings     int
	NotesWithMissingFiles int
	TotalWarnings         int
	TotalMissingFiles     int
	WarningTypes          map[string]int
	MissingFileTypes      map[string]int
}

type DiagnosticsFormatter struct {
	red       lipgloss.Color
	orange    lipgloss.Color
	yellow    lipgloss.Color
	gray      lipgloss.Color
	lightGray lipgloss.Color

	headerStyle  lipgloss.Style
	cellStyle    lipgloss.Style
	oddRowStyle  lipgloss.Style
	evenRowStyle lipgloss.Style
}

func NewDiagnosticsFormatter() *DiagnosticsFormatter {
	red := lipgloss.Color("#b64229")
	orange := lipgloss.Color("208")
	yellow := lipgloss.Color("#b69c29")
	gray := lipgloss.Color("245")
	lightGray := lipgloss.Color("241")

	headerStyle := lipgloss.NewStyle().Foreground(red).Bold(true).Align(lipgloss.Center)
	cellStyle := lipgloss.NewStyle().Padding(1, 1).Width(22)
	oddRowStyle := cellStyle.Foreground(gray)
	evenRowStyle := cellStyle.Foreground(lightGray)

	return &DiagnosticsFormatter{
		red:          red,
		orange:       orange,
		yellow:       yellow,
		gray:         gray,
		lightGray:    lightGray,
		headerStyle:  headerStyle,
		cellStyle:    cellStyle,
		oddRowStyle:  oddRowStyle,
		evenRowStyle: evenRowStyle,
	}
}

func (df *DiagnosticsFormatter) FormatDiagnosticsSummary(summary DiagnosticsSummary) string {
	if summary.TotalWarnings == 0 && summary.TotalMissingFiles == 0 {
		return "\n✅ No parsing issues detected - vault is healthy!"
	}

	var output strings.Builder
	output.WriteString("\n🔍 VAULT DIAGNOSTICS SUMMARY\n")

	mainTable := df.createMainTable(summary)
	output.WriteString(mainTable)
	output.WriteString("\n")

	if summary.TotalWarnings > 0 && len(summary.WarningTypes) > 0 {
		warningTable := df.createWarningTypesTable(summary.WarningTypes, summary.TotalWarnings)
		output.WriteString(warningTable)
		output.WriteString("\n")
	}

	if summary.TotalMissingFiles > 0 && len(summary.MissingFileTypes) > 0 {
		missingTable := df.createMissingFileTypesTable(summary.MissingFileTypes, summary.TotalMissingFiles)
		output.WriteString(missingTable)
		output.WriteString("\n")
	}

	return output.String()
}

func (df *DiagnosticsFormatter) createMainTable(summary DiagnosticsSummary) string {
	mainTable := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(df.red)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return df.headerStyle
			case row%2 == 0:
				return df.evenRowStyle
			default:
				return df.oddRowStyle
			}
		}).
		Headers("METRIC", "COUNT", "PERCENTAGE")

	if summary.TotalWarnings > 0 {
		warningPercentage := fmt.Sprintf("%.1f%%", float64(summary.NotesWithWarnings)/float64(summary.TotalNotes)*100)
		mainTable.Row("⚠️  Total Warnings", fmt.Sprintf("%d", summary.TotalWarnings), "-")
		mainTable.Row("📄 Notes w/ Warnings", fmt.Sprintf("%d", summary.NotesWithWarnings), warningPercentage)
	}

	if summary.TotalMissingFiles > 0 {
		missingPercentage := fmt.Sprintf("%.1f%%", float64(summary.NotesWithMissingFiles)/float64(summary.TotalNotes)*100)
		mainTable.Row("❌ Missing Files", fmt.Sprintf("%d", summary.TotalMissingFiles), "-")
		mainTable.Row("📄 Notes w/ Missing", fmt.Sprintf("%d", summary.NotesWithMissingFiles), missingPercentage)
	}

	mainTable.Row("📊 Total Notes", fmt.Sprintf("%d", summary.TotalNotes), "100%")

	return mainTable.String()
}

func (df *DiagnosticsFormatter) createWarningTypesTable(warningTypes map[string]int, totalWarnings int) string {
	var output strings.Builder
	output.WriteString("\n⚠️  WARNING TYPES BREAKDOWN\n")

	warningTable := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(df.orange)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Foreground(df.orange).Bold(true).Align(lipgloss.Center)
			case row%2 == 0:
				return df.evenRowStyle
			default:
				return df.oddRowStyle
			}
		}).
		Headers("WARNING TYPE", "COUNT", "PERCENTAGE")

	for warningType, count := range warningTypes {
		percentage := fmt.Sprintf("%.1f%%", float64(count)/float64(totalWarnings)*100)
		var displayType string
		switch warningType {
		case "unknown_frontmatter":
			displayType = "Unknown Frontmatter"
		case "frontmatter_parsing":
			displayType = "Frontmatter Parsing"
		default:
			temp := strings.ReplaceAll(warningType, "_", " ")
			displayType = strings.ToUpper(temp[:1]) + temp[1:]
		}
		warningTable.Row(displayType, fmt.Sprintf("%d", count), percentage)
	}

	output.WriteString(warningTable.String())
	return output.String()
}

func (df *DiagnosticsFormatter) createMissingFileTypesTable(missingFileTypes map[string]int, totalMissingFiles int) string {
	var output strings.Builder
	output.WriteString("\n📁 MISSING FILE TYPES BREAKDOWN\n")

	missingTable := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(df.yellow)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Foreground(df.yellow).Bold(true).Align(lipgloss.Center)
			case row%2 == 0:
				return df.evenRowStyle
			default:
				return df.oddRowStyle
			}
		}).
		Headers("FILE TYPE", "COUNT", "PERCENTAGE")

	for fileType, count := range missingFileTypes {
		percentage := fmt.Sprintf("%.1f%%", float64(count)/float64(totalMissingFiles)*100)
		displayType := fileType
		if fileType == "no_extension" {
			displayType = "No Extension"
		}
		missingTable.Row(displayType, fmt.Sprintf("%d", count), percentage)
	}

	output.WriteString(missingTable.String())
	return output.String()
}
