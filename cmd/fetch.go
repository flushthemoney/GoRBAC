package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/flushthemoney/GoRBAC/internal/k8s"
	"github.com/flushthemoney/GoRBAC/internal/types"
	"github.com/spf13/cobra"
)

var kubeconfig string
var namespaces string
var jsonOut bool

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch RBAC resources from a Kubernetes cluster.",
	Long: `Fetches Roles, ClusterRoles, RoleBindings, and ClusterRoleBindings from a Kubernetes cluster.
You can save the results to a JSON file for further analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := fetchRBAC(kubeconfig, namespaces, jsonOut)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	fetchCmd.Flags().StringVar(&namespaces, "namespace", "", "Namespaces to audit")
	fetchCmd.Flags().BoolVar(&jsonOut, "json-out", false, "Check to save RBAC details to JSON")
}

func fetchRBAC(kubeconfig string, namespace string, jsonOut bool) error {
	clientset, err := k8s.NewClient(kubeconfig)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	resources, err := clientset.GetRBACResources(context.Background(), namespace)
	if err != nil {
		return fmt.Errorf("failed to get RBAC resources: %w", err)
	}

	if jsonOut {
		// Populate metadata
		meta := types.Metadata{
			ClusterName: "my-cluster", // Placeholder
			Timestamp:   time.Now(),
		}
		if namespace != "" {
			meta.Namespaces = strings.Split(namespace, ",")
		}

		resources.Metadata = meta

		jsonData, err := json.MarshalIndent(resources, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal data to JSON: %w", err)
		}
		filename := "rbac_resources.json"
		err = os.WriteFile(filename, jsonData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write json file: %w", err)
		}

		fmt.Printf("Successfully wrote wrote JSON to %s\n", filename)
	}

	return nil
}
