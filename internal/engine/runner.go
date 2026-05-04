package engine

import (
	"github.com/gintorust/clone-sage/internal/model"
	"github.com/gintorust/clone-sage/internal/registry"
)

//Run executes the planned checks and aggregates the results
func Run(configs []model.CheckConfig, timeoutMs int) ([]model.Result, error){
	var results []model.Result

	for _, cfg := range configs {
		//build check from registry
		check, err := registry.Build(cfg)
		if err != nil {
			return nil, err
		}

		res:= RunWithTimeout(check, timeoutMs)

		//we keep the original severity provided by the user or the system
		res.Severity = cfg.Severity
		results = append(results, res)
	}

	return results, nil
}