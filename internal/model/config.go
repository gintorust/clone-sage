package model

// CloneSageConfig represents the top-level configuration file.
type CloneSageConfig struct {
	Version  int                    `yaml:"version"`
	Defaults map[string]interface{} `yaml:"defaults"`
	Checks   []CheckConfig          `yaml:"checks"`
}

// CheckConfig represents a single diagnostic check defined in YAML.
type CheckConfig struct {
	Name     string            `yaml:"name"`
	Type     string            `yaml:"type"`
	Quick    bool              `yaml:"quick"`
	Severity string            `yaml:"severity"`
	Options  map[string]string `yaml:"options"`
	Message  string            `yaml:"message"`
	Why      string            `yaml:"why,omitempty"`
	Fix      string            `yaml:"fix,omitempty"`
}
