package detect

import "github.com/gintorust/clone-sage/internal/model"

func detectGo() []model.CheckConfig {
	var checks []model.CheckConfig

	if fileExists("go.mod") {
		checks = append(checks, model.CheckConfig{
			Name:     "go-installed",
			Type:     "command_exists",
			Severity: "blocker",
			Options:  map[string]string{"command": "go"},
			Message:  "Go is required for this repository",
		})
	}

	return checks
}
