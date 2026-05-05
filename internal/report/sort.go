package report

import (
	"sort"

	"github.com/gintorust/clone-sage/internal/model"
)

func severityWeight(severity string) int{
	switch severity{
	case "info":
		return 1
	case "warning":
		return 2
	case "blocker":
		return 3
	}
	return 0
}

// SortBySeverity sorts results: Failed Blockers > Failed Warnings > Passed
func SortBySeverity(results []model.Result) {
	sort.Slice(results, func(i, j int) bool {
		// First sort by status (failed bubbles up)
		if results[i].Status != results[j].Status {
			return results[i].Status == "failed"
		}
		// Then sort by severity weight
		return severityWeight(results[i].Severity) > severityWeight(results[j].Severity)
	})
}