package detect

// Looks for Node signatures and returns relevant checks
func detectNode() []CheckConfig {
	var checks []CheckConfig

	if fileExists("package.json") {
		checks = append(checks, CheckConfig{
			Name:     "node-installed",
			Type:     "command_exists",
			Severity: "blocker",
			Options:  map[string]string{"command": "node"},
			Message:  "Node.js is required to run this project.",
		})
		
	}

	return checks
}