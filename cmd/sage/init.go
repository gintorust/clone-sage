/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package sage

import (
	"fmt"
	"os"

	"github.com/gintorust/clone-sage/internal/detect"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type CloneSageConfig struct{
	Version int                     `yaml:"version"`
	Defaults map[string]interface{} `yaml:"defaults"`
    Checks []detect.CheckConfig     `yaml:"checks"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a clonesage-auto.yaml file based on repository inference",
	Long: `Scans the current directory for project files (go.mod, package.json, .env.example) 
and automatically generates a tailored clonesage-auto.yaml configuration file.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		_,err := os.Stat("clonesage-auto.yaml")
		if err == nil {
			return fmt.Errorf("clonesage-auto.yaml already exists. Please delete or rename the existing file before running 'sage init'")
		}

		fmt.Println("Scanning repo...")

		discoveredChecks := detect.ScanRepo()

		config := CloneSageConfig{
			Version: 1,
			Defaults: map[string]interface{}{
				"strict": true,
				"timeout_ms": 3000,
			},
			Checks: discoveredChecks,
		}

		yamlData, err := yaml.Marshal(&config)
		if err != nil {
			return fmt.Errorf("Failed to generate yaml file: %w", err)
		}

		err = os.WriteFile("clonesage-auto.yaml", yamlData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write clonesage.yaml: %w", err)
		}

		fmt.Printf("Success! Generated clonesage.yaml with %d inferred checks.\n", len(discoveredChecks))
		return nil

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
