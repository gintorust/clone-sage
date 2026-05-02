package detect

import "github.com/gintorust/clone-sage/internal/model"

// Looks for Node signatures and returns relevant checks
func detectNode() []model.CheckConfig {
	var checks []model.CheckConfig

	if fileExists("package.json") {
		checks = append(checks, model.CheckConfig{
			Name:     "node-installed",
			Type:     "command_exists",
			Severity: "blocker",
			Options:  map[string]string{"command": "node"},
			Message:  "Node.js is required to run this project.",
		})
		
	}

	return checks
}