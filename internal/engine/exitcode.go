package engine

import "github.com/gintorust/clone-sage/internal/model"

func CalculateExitCode(results []model.Result) int{
	for _, result := range results{
		if result.Status == "failed" && result.Severity == "blocker"{
			return 1 //blockers found
		}
	}
	return 0 //clean or only warnings
}