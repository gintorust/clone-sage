package detect

func detectDocker() []CheckConfig {
	var checks []CheckConfig

	if fileExists("docker-compose.yml") || fileExists("compose.yaml") || fileExists("docker-compose.yaml") {
		checks = append(checks, CheckConfig{
			Name:     "docker-installed",
			Type:     "command_exists",
			Severity: "warning", 
			Options:  map[string]string{"command": "docker"},
			Message:  "A Docker Compose file was detected. Docker is recommended for local infrastructure.",
			Fix:      "Install Docker",
		})
	}

	return checks
}