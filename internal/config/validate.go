package config

import (
	"fmt"
	"strings"

	"github.com/gintorust/clone-sage/internal/model"
)

func ValidateConfig(c *model.CloneSageConfig) error {
	var errors []string

	if c.Version != 1 {
		errors = append(errors, fmt.Sprintf("-unsupported version: %d (expected 1)", c.Version))
	}

	seenNames := make(map[string]bool)

	for i, check := range c.Checks {
		prefix := fmt.Sprintf("checks[%d] ('%s')", i, check.Name)

		if check.Name == "" {
			errors = append(errors, fmt.Sprintf("%s: name is required", prefix))
		} else if seenNames[check.Name] {
			errors = append(errors, fmt.Sprintf("%s: duplicate check name found", prefix))
		}
		seenNames[check.Name] = true

		if check.Severity != "info" && check.Severity != "warning" && check.Severity != "blocker" {
			errors = append(errors, fmt.Sprintf("%s: invalid severity '%s' (must be info, warning, or blocker)", prefix, check.Severity))
		}

		if check.Type == "" {
			errors = append(errors, fmt.Sprintf("%s: type is required", prefix))
		}

		if check.Message == "" {
			errors = append(errors, fmt.Sprintf("%s: message is required to explain failures", prefix))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("configuration validation failed:\n%s", strings.Join(errors, "\n"))
	}

	return nil

}
