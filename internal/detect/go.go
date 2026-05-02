package detect

func detectGo() []CheckConfig {
	var checks []CheckConfig

	if fileExists("go.mod") {
		checks = append(checks, CheckConfig{
			Name:     "go-installed",
			Type:     "command-exists",
			Severity: "blocker",
			Options:  map[string]string{"command": "go"},
			Message:  "Go is required for this repository",
		})
	}

	return checks
}
