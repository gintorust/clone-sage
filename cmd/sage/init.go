/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package sage

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gintorust/clone-sage/internal/detect"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type CloneSageConfig struct {
	Version  int                    `yaml:"version"`
	Defaults map[string]interface{} `yaml:"defaults"`
	Checks   []detect.CheckConfig   `yaml:"checks"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a sage-auto.yaml file based on repository inference",
	Long: `Scans the current directory for project files (go.mod, package.json, .env.example) 
and automatically generates a tailored sage-auto.yaml configuration file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := os.Stat("sage-auto.yaml")
		if err == nil {
			return fmt.Errorf("sage-auto.yaml already exists. Please delete or rename the existing file before running 'sage init'")
		}

		fmt.Println("Scanning repo...")

		discoveredChecks := detect.ScanRepo()

		config := CloneSageConfig{
			Version: 1,
			Defaults: map[string]interface{}{
				"strict":     true,
				"timeout_ms": 3000,
			},
			Checks: discoveredChecks,
		}

		yamlData, err := yaml.Marshal(&config)
		if err != nil {
			return fmt.Errorf("Failed to generate yaml file: %w", err)
		}

		yamlString := string(yamlData)

		yamlString = strings.Replace(yamlString, "\nchecks:", "\n\nchecks:", 1)

		re := regexp.MustCompile(`\n(\s*- name:)`)
		yamlString = re.ReplaceAllString(yamlString, "\n\n$1")
		
		yamlString = strings.ReplaceAll(yamlString, "\n\n\n", "\n\n")

		finalYamlData := []byte(yamlString)

		err = os.WriteFile("sage-auto.yaml", finalYamlData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write sage.yaml: %w", err)
		}

		fmt.Printf("Success! Generated sage-auto.yaml with %d inferred checks.\n", len(discoveredChecks))
		return nil

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
