package terminal

import (
	"fmt"
	"strings"
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

func (pp *PrettyPrinter) PrintDiagnostics(summary DiagnosticsSummary) string {
	if summary.TotalWarnings == 0 && summary.TotalMissingFiles == 0 {
		return pp.Success("No parsing issues detected - vault is healthy!")
	}

	var output strings.Builder

	output.WriteString("\n\n")

	output.WriteString(pp.formatDiagnosticMetrics(summary))
	output.WriteString("\n")

	if summary.TotalWarnings > 0 && len(summary.WarningTypes) > 0 {
		output.WriteString(pp.formatWarningBreakdown(summary.WarningTypes, summary.TotalWarnings))
		output.WriteString("\n")
	}

	if summary.TotalMissingFiles > 0 && len(summary.MissingFileTypes) > 0 {
		output.WriteString(pp.formatMissingFilesBreakdown(summary.MissingFileTypes, summary.TotalMissingFiles))
		output.WriteString("\n")
	}

	return output.String()
}

func (pp *PrettyPrinter) formatDiagnosticMetrics(summary DiagnosticsSummary) string {
	var output strings.Builder

	output.WriteString(pp.Info("Total Notes", "count", summary.TotalNotes))
	output.WriteString("\n")

	if summary.TotalWarnings > 0 {
		warningPercentage := fmt.Sprintf("%.1f%%", float64(summary.NotesWithWarnings)/float64(summary.TotalNotes)*100)
		output.WriteString(pp.Warning("Notes with warnings", "count", summary.NotesWithWarnings, "percentage", warningPercentage))
		output.WriteString("\n")

		output.WriteString(pp.Warning("Total warnings", "count", summary.TotalWarnings))
		output.WriteString("\n")
	}

	if summary.TotalMissingFiles > 0 {
		missingPercentage := fmt.Sprintf("%.1f%%", float64(summary.NotesWithMissingFiles)/float64(summary.TotalNotes)*100)
		output.WriteString(pp.Error("Notes with missing files", "count", summary.NotesWithMissingFiles, "percentage", missingPercentage))
		output.WriteString("\n")

		output.WriteString(pp.Error("Total missing files", "count", summary.TotalMissingFiles))
		output.WriteString("\n")
	}

	return output.String()
}

func (pp *PrettyPrinter) formatWarningBreakdown(warningTypes map[string]int, totalWarnings int) string {
	var output strings.Builder

	output.WriteString(pp.Warning("Warning Types Breakdown"))
	output.WriteString("\n")

	for warningType, count := range warningTypes {
		percentage := fmt.Sprintf("%.1f%%", float64(count)/float64(totalWarnings)*100)
		displayType := pp.formatWarningType(warningType)

		line := fmt.Sprintf("  • %s", displayType)
		output.WriteString(pp.formatMessage(line, "count", count, "percentage", percentage))
		output.WriteString("\n")
	}

	return output.String()
}

func (pp *PrettyPrinter) formatMissingFilesBreakdown(missingFileTypes map[string]int, totalMissingFiles int) string {
	var output strings.Builder

	output.WriteString(pp.Error("❌ Missing File Types Breakdown"))
	output.WriteString("\n")

	for fileType, count := range missingFileTypes {
		percentage := fmt.Sprintf("%.1f%%", float64(count)/float64(totalMissingFiles)*100)
		displayType := fileType
		if fileType == "no_extension" {
			displayType = "No Extension"
		}

		line := fmt.Sprintf("  • %s", displayType)
		output.WriteString(pp.formatMessage(line, "count", count, "percentage", percentage))
		output.WriteString("\n")
	}

	return output.String()
}

func (pp *PrettyPrinter) formatWarningType(warningType string) string {
	switch warningType {
	case "unknown_frontmatter":
		return "Unknown Frontmatter"
	case "frontmatter_parsing":
		return "Frontmatter Parsing"
	default:
		temp := strings.ReplaceAll(warningType, "_", " ")
		return strings.ToUpper(temp[:1]) + temp[1:]
	}
}

func (pp *PrettyPrinter) PrintDiagnosticsSummary(summary DiagnosticsSummary) string {
	if summary.TotalWarnings == 0 && summary.TotalMissingFiles == 0 {
		return pp.Success("Vault is healthy", "notes", summary.TotalNotes)
	}

	var issues []string
	if summary.TotalWarnings > 0 {
		issues = append(issues, fmt.Sprintf("warnings=%d", summary.TotalWarnings))
	}
	if summary.TotalMissingFiles > 0 {
		issues = append(issues, fmt.Sprintf("missing=%d", summary.TotalMissingFiles))
	}

	message := "Vault diagnostics"
	keyvals := []any{"notes", summary.TotalNotes}
	for _, issue := range issues {
		parts := strings.SplitN(issue, "=", 2)
		if len(parts) == 2 {
			keyvals = append(keyvals, parts[0], parts[1])
		}
	}

	return pp.Warning(message, keyvals...)
}
