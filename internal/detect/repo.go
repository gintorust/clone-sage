package detect

import "strings"


type CheckConfig struct {
	Name string                `yaml:"name"`
	Type string                `yaml:"type"`
	Severity string            `yaml:"severity"`
	Options map[string]string  `yaml:"options"`
	Message string             `yaml:"message"`
	Fix string                 `yaml:"fix"`
}

func ScanRepo() []CheckConfig{
	var allChecks []CheckConfig

	allChecks = append(allChecks, detectGo()...)
	allChecks = append(allChecks, detectNode()...)
	allChecks = append(allChecks, detectDocker()...)

	if fileExists("Makefile") {
		allChecks = append(allChecks, CheckConfig{
			Name:     "make-installed",
			Type:     "command_exists",
			Severity: "blocker",
			Options:  map[string]string{"command": "make"},
			Message:  "A Makefile was detected. 'make' is required to run build tasks.",
		})
	}

	if fileExists("Taskfile.yml") || fileExists("Taskfile.yaml") {
		allChecks = append(allChecks, CheckConfig{
			Name:     "task-installed",
			Type:     "command_exists",
			Severity: "blocker",
			Options:  map[string]string{"command": "task"},
			Message:  "A Taskfile was detected. The 'task' runner is required.",
			Fix:      "Run 'brew install go-task/tap/go-task'",
		})
	}

	if fileExists(".tool-versions") {
		allChecks = append(allChecks, CheckConfig{
			Name:     "asdf-installed",
			Type:     "command_exists",
			Severity: "warning",
			Options:  map[string]string{"command": "asdf"},
			Message:  ".tool-versions detected. A version manager like asdf or mise is required.",
		})
	}

	if fileExists(".env.example"){
		allChecks = append(allChecks, CheckConfig{
			Name:     "env-file-exists",
			Type:     "file_exists",
			Severity: "blocker",
			Options:  map[string]string{"path": ".env"},
			Message:  "You must create a .env file.",
			Fix:      "Run 'cp .env.example .env'",
		})

		keys := extractEnvKeys(".env.example")
		for _, key := range keys {
			allChecks = append(allChecks, CheckConfig{
				Name:     strings.ToLower(key) + "-configured",
				Type:     "env_exists",
				Severity: "blocker",
				Options:  map[string]string{"key": key},
				Message:  key + " is missing from your environment.",
			})
		}

	}

	return allChecks
}