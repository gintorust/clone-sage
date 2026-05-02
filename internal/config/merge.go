package config

import "github.com/gintorust/clone-sage/internal/model"

func mergeDefaults(c *model.CloneSageConfig) {
	if c.Defaults == nil {
		c.Defaults = make(map[string]interface{})
	}

	if _, exists := c.Defaults["strict"]; !exists {
		c.Defaults["strict"] = true
	}

	if _, exists := c.Defaults["timeout_ms"]; !exists {
		c.Defaults["timeout_ms"] = 3000
	}

}
