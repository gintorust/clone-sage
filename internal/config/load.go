package config

import (
	"fmt"
	"os"

	"github.com/gintorust/clone-sage/internal/model"
	"github.com/spf13/viper"
)

//this function is used for handling the auto-generated and the explicit yaml files by parsing, merging and validating configurations
func Load(cfgFile string) (*model.CloneSageConfig, error){
	autoCfg, errAuto := readConfigFile("sage-auto.yaml")

	explicitFileName := "sage.yaml"
	if cfgFile != ""{
		explicitFileName = cfgFile
	}

	explicitCfg, errExplicit := readConfigFile(explicitFileName)

	if os.IsNotExist(errAuto) && os.IsNotExist(errExplicit){
		return nil, fmt.Errorf("No config files found. Run 'sage init' to generate config files")
	}

	var finalCfg *model.CloneSageConfig

	if explicitCfg == nil {
		fmt.Println("Using auto-generated config (no sage.yaml found).")
		finalCfg = autoCfg
	} else if autoCfg == nil {
		fmt.Println("Using explicit config (no clonesage-auto.yaml found).")
		finalCfg = explicitCfg
	} else {
		fmt.Println("Merging sage-auto.yaml with sage.yaml...")
		finalCfg = mergeConfigs(autoCfg, explicitCfg)
	}

	MergeDefaults(finalCfg)

	if err := ValidateConfig(finalCfg); err != nil{
		return nil, fmt.Errorf("Invalid merged configurations: \n%s", err)
	}

	return finalCfg, nil
}

//this function is used for reading the given yaml and unmarshaling it into a CloneSageConfig struct
func readConfigFile(filename string) (*model.CloneSageConfig, error){
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(filename)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read %s: %w", filename, err)
	}

	var cfg model.CloneSageConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Failed to parse %s: %w", filename, err)
	}

	return &cfg, nil
}

// mergeConfigs safely layers the explicit config over the auto config
func mergeConfigs(auto, explicit *model.CloneSageConfig) *model.CloneSageConfig {
	merged := &model.CloneSageConfig{
		Version:  explicit.Version,
		Defaults: make(map[string]interface{}),
	}
	if merged.Version == 0 {
		merged.Version = auto.Version // Fallback if explicit didn't define it
	}

	for k, v := range auto.Defaults {
		merged.Defaults[k] = v
	}
	for k, v := range explicit.Defaults {
		merged.Defaults[k] = v
	}

	checkMap := make(map[string]model.CheckConfig)
	var order []string

	for _, c := range auto.Checks {
		checkMap[c.Name] = c
		order = append(order, c.Name)
	}

	for _, c := range explicit.Checks {
		if _, exists := checkMap[c.Name]; !exists {
			order = append(order, c.Name)
		}
		checkMap[c.Name] = c
	}

	for _, name := range order {
		merged.Checks = append(merged.Checks, checkMap[name])
	}

	return merged
}