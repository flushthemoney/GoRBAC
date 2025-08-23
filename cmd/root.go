/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
// save them to a JSON file for further analysis.`,
var rootCmd = &cobra.Command{
	Use:   "rbaclens",
	Short: "RBACLens: Audit and analyze Kubernetes RBAC resources.",
	Long: `RBACLens is a CLI tool designed to help you audit and analyze
Kubernetes RBAC (Role-Based Access Control) resources.

Features:
- Fetch RBAC resources from a cluster and save them to a JSON file.
- Audit RBAC resources for risky configurations using built-in rules.

See the documentation for details on each command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rbaclens.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
