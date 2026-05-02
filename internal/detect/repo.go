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