package report

import (
	"fmt"

	"github.com/gintorust/clone-sage/internal/model"
)

// PrintTerminal formats the results for a human reader.
func PrintTerminal(results []model.Result) {
	SortBySeverity(results)

	var blockers, warnings, passed int

	fmt.Println("\n--- CloneSage Diagnostics ---")

	for _, res := range results {
		if res.Status == "passed" {
			passed++
			continue // We don't spam the terminal with passed checks
		}

		if res.Severity == "blocker" {
			blockers++
			fmt.Printf("\n[BLOCKER] %s\n", res.Name)
		} else {
			warnings++
			fmt.Printf("\n[WARNING] %s\n", res.Name)
		}

		fmt.Printf("  Evidence: %s\n", res.Evidence)
		if res.Why != "" {
			fmt.Printf("  Why:      %s\n", res.Why)
		}
		if res.Fix != "" {
			fmt.Printf("  Fix:      %s\n", res.Fix)
		}
	}

	// Print the summary at the end
	fmt.Printf("\nSummary: %d blockers, %d warnings, %d passed\n\n", blockers, warnings, passed)
}
