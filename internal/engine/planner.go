package engine

import "github.com/gintorust/clone-sage/internal/model"

func Plan(configs []model.CheckConfig, isQuickMode bool) []model.CheckConfig {
	var planned []model.CheckConfig

	for _, cfg := range configs {
		if isQuickMode && !cfg.Quick {
			continue
		}
		planned = append(planned, cfg)
	}

	return planned
}
