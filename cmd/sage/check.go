/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package sage

import (
	"fmt"
	"os"

	"github.com/gintorust/clone-sage/internal/config"
	"github.com/gintorust/clone-sage/internal/engine"
	"github.com/gintorust/clone-sage/internal/report"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	isQuickMode bool
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run the diagnostic suite for the current repository",
	Long: `Parses the clonesage.yaml file and executes all defined checks 
to diagnose local development setup failures.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			os.Exit(2) // Exit 2 for Configuration Errors
		}

		// Plan
		planned := engine.Plan(cfg.Checks, isQuickMode)

		// Run
		results, err := engine.Run(planned, cfg.Defaults["timeout_ms"].(int))
		if err != nil {
			fmt.Printf("Internal Execution Error: %v\n", err)
			os.Exit(3) // Exit 3 for Internal Errors
		}

		// Report
		report.PrintTerminal(results)

		// Exit
		os.Exit(engine.CalculateExitCode(results))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
