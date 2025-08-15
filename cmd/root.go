/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gorbac",
	Short: "GoRBAC is a CLI tool for auditing Kubernetes RBAC resources.",
	Long: `GoRBAC is a CLI tool designed to help you audit and analyze
Kubernetes RBAC (Role-Based Access Control) resources. 

You can use GoRBAC to fetch RBAC resources from a cluster and 
save them to a JSON file for further analysis.`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gorbac.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}


