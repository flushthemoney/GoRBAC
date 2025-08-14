package cmd

import (
	"fmt"
	"log"

	"github.com/flushthemoney/GoRBAC/internal/utils"
	"github.com/spf13/cobra"
)

var kubeconfig string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := fetchRBAC(kubeconfig)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")

}

func fetchRBAC(kubeconfig string) error {
	_, err := utils.NewClient(kubeconfig)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}
	return nil
}
