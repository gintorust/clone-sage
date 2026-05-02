package detect

import "github.com/gintorust/clone-sage/internal/model"

func detectDocker() []model.CheckConfig {
	var checks []model.CheckConfig

	if fileExists("docker-compose.yml") || fileExists("compose.yaml") || fileExists("docker-compose.yaml") {
		checks = append(checks, model.CheckConfig{
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