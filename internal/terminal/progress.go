package terminal

import (
	"fmt"
	"time"
)

type ProgressReporter struct {
	printer *PrettyPrinter
}

func NewProgressReporter(printer *PrettyPrinter) *ProgressReporter {
	return &ProgressReporter{
		printer: printer,
	}
}

type Step struct {
	Name string
	Icon string
}

type StepResult struct {
	Duration time.Duration
	Count    int
	Error    error
}

func (pr *ProgressReporter) ReportStep(step Step, result StepResult) {
	if result.Error != nil {
		message := fmt.Sprintf("%s %s failed", step.Icon, step.Name)
		fmt.Println(pr.printer.Error(message, "error", result.Error))
		return
	}

	message := fmt.Sprintf("%s %s", step.Icon, step.Name)
	fmt.Println(pr.printer.Print(message, "count", result.Count, "duration", result.Duration))
}

func (pr *ProgressReporter) ReportStepInProgress(step Step) {
	message := fmt.Sprintf("%s %s", step.Icon, step.Name)
	fmt.Println(pr.printer.Info(message))
}
