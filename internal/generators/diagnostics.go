package generators

import (
	"encoding/json"
	"fmt"
	"gobsidian/internal/models"
	"gobsidian/internal/terminal"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (g *StaticSiteGenerator) printDiagnosticsSummary(notes []*models.ParsedNote, printer *terminal.PrettyPrinter) {
	var totalWarnings, totalMissingFiles int
	var notesWithWarnings, notesWithMissingFiles int
	var warningTypes = make(map[string]int)
	var missingFileTypes = make(map[string]int)

	for _, note := range notes {
		if note.HasWarnings() {
			notesWithWarnings++
			totalWarnings += len(note.Warnings)

			for _, warning := range note.Warnings {
				if strings.Contains(warning, "unknown frontmatter field") {
					warningTypes["unknown_frontmatter"]++
				} else if strings.Contains(warning, "frontmatter parsing issues") {
					warningTypes["frontmatter_parsing"]++
				} else {
					warningTypes["other"]++
				}
			}
		}

		if note.HasMissingFiles() {
			notesWithMissingFiles++
			totalMissingFiles += len(note.MissingFiles)

			for _, file := range note.MissingFiles {
				ext := strings.ToLower(filepath.Ext(file))
				if ext == "" {
					ext = "no_extension"
				}
				missingFileTypes[ext]++
			}
		}
	}

	summary := terminal.DiagnosticsSummary{
		TotalNotes:            len(notes),
		NotesWithWarnings:     notesWithWarnings,
		NotesWithMissingFiles: notesWithMissingFiles,
		TotalWarnings:         totalWarnings,
		TotalMissingFiles:     totalMissingFiles,
		WarningTypes:          warningTypes,
		MissingFileTypes:      missingFileTypes,
	}

	output := printer.PrintDiagnostics(summary)
	fmt.Println(printer.Box("Vault diagnostics", output))

	if totalWarnings > 0 || totalMissingFiles > 0 {
		g.exportDiagnostics(notes, warningTypes, missingFileTypes, totalWarnings, totalMissingFiles)
	}
}

func (g *StaticSiteGenerator) exportDiagnostics(
	notes []*models.ParsedNote,
	warningTypes map[string]int,
	missingFileTypes map[string]int,
	totalWarnings int,
	totalMissingFiles int,
) {
	type NoteDiagnostic struct {
		Title        string   `json:"title"`
		RelativePath string   `json:"relative_path"`
		Warnings     []string `json:"warnings"`
		MissingFiles []string `json:"missing_files"`
		TotalIssues  int      `json:"total_issues"`
	}

	type DiagnosticReport struct {
		GeneratedAt           string           `json:"generated_at"`
		TotalNotes            int              `json:"total_notes"`
		NotesWithWarnings     int              `json:"notes_with_warnings"`
		NotesWithMissingFiles int              `json:"notes_with_missing_files"`
		TotalWarnings         int              `json:"total_warnings"`
		TotalMissingFiles     int              `json:"total_missing_files"`
		WarningTypes          map[string]int   `json:"warning_types"`
		MissingFileTypes      map[string]int   `json:"missing_file_types"`
		ProblematicNotes      []NoteDiagnostic `json:"problematic_notes"`
	}

	if totalWarnings == 0 && totalMissingFiles == 0 {
		return
	}

	var problematicNotes []NoteDiagnostic
	var notesWithWarnings, notesWithMissingFiles int

	for _, note := range notes {
		if note.HasWarnings() || note.HasMissingFiles() {
			totalIssues := len(note.Warnings) + len(note.MissingFiles)
			problematicNotes = append(problematicNotes, NoteDiagnostic{
				Title:        note.Title,
				RelativePath: note.RelativePath,
				Warnings:     note.Warnings,
				MissingFiles: note.MissingFiles,
				TotalIssues:  totalIssues,
			})

			if note.HasWarnings() {
				notesWithWarnings++
			}
			if note.HasMissingFiles() {
				notesWithMissingFiles++
			}
		}
	}

	report := DiagnosticReport{
		GeneratedAt:           time.Now().Format(time.RFC3339),
		TotalNotes:            len(notes),
		NotesWithWarnings:     notesWithWarnings,
		NotesWithMissingFiles: notesWithMissingFiles,
		TotalWarnings:         totalWarnings,
		TotalMissingFiles:     totalMissingFiles,
		WarningTypes:          warningTypes,
		MissingFileTypes:      missingFileTypes,
		ProblematicNotes:      problematicNotes,
	}

	diagnosticsPath := filepath.Join(g.SiteConfig.OutputDirectory, diagnosticsFile)
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		g.Logger.Warn("Failed to marshal diagnostics report", "error", err)
		return
	}

	if err := os.WriteFile(diagnosticsPath, jsonData, 0644); err != nil {
		g.Logger.Warn("Failed to write diagnostics report", "error", err)
		return
	}

	g.Logger.Debug("📊 Diagnostics report exported",
		"path", diagnosticsPath,
		"size", fmt.Sprintf("%.1f KB", float64(len(jsonData))/1024),
	)
}
