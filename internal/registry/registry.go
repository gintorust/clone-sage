package registry

import (
	"fmt"

	"github.com/gintorust/clone-sage/internal/model"
)

// CheckFactory is a function that takes the YAML config and returns an executable Check.
type CheckFactory func(cfg model.CheckConfig) model.Check

var checkRegistry = make(map[string]CheckFactory)

// Register is called by individual check files (e.g., in their init() functions)
func Register(checkType string, factory CheckFactory) {
	checkRegistry[checkType] = factory
}

// Build looks up the type and constructs the check.
func Build(cfg model.CheckConfig) (model.Check, error) {
	factory, exists := checkRegistry[cfg.Type]
	if !exists {
		return nil, fmt.Errorf("unknown check type: %s", cfg.Type)
	}
	return factory(cfg), nil
}
