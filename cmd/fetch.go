package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/flushthemoney/GoRBAC/internal/utils"
	"github.com/spf13/cobra"
)

var kubeconfig string
var namespaces string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := fetchRBAC(kubeconfig, namespaces)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	fetchCmd.Flags().StringVar(&namespaces, "namespace", "", "Namespaces to audit")

}

func fetchRBAC(kubeconfig string, namespace string) error {
	clientset, err := utils.NewClient(kubeconfig)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	_, err = clientset.GetRBACResources(context.Background(), namespace)
	if err != nil {
		return fmt.Errorf("failed to get RBAC resources: %w", err)
	}

	return nil
}
