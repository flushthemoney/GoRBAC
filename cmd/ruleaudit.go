package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/flushthemoney/GoRBAC/internal/audit"
	"github.com/flushthemoney/GoRBAC/internal/k8s"
	"github.com/flushthemoney/GoRBAC/internal/types"
	"github.com/spf13/cobra"
)

// Only declare variables not already declared in fetch.go
var namespace string
var inputFile string

// ruleAuditCmd represents the ruleaudit command
var ruleAuditCmd = &cobra.Command{
	Use:   "ruleaudit",
	Short: "Audit RBAC resources for risky configurations",
	Long: `Audit RBAC resources for risky configurations using built-in rules.
You can fetch live from a cluster or audit a previously saved JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var resources types.RBACResources
		if inputFile != "" {
			data, err := os.ReadFile(inputFile)
			if err != nil {
				log.Fatalf("Failed to read input file: %v", err)
			}
			if err := json.Unmarshal(data, &resources); err != nil {
				log.Fatalf("Failed to unmarshal input file: %v", err)
			}
		} else {
			clientset, err := k8s.NewClient(kubeconfig)
			if err != nil {
				log.Fatalf("Failed to create Kubernetes client: %v", err)
			}
			rsrcPtr, err := clientset.GetRBACResources(cmd.Context(), namespace)
			if err != nil {
				log.Fatalf("Failed to get RBAC resources: %v", err)
			}
			resources = *rsrcPtr
			meta := types.Metadata{
				ClusterName: "my-cluster", // Placeholder
				Timestamp:   time.Now(),
			}
			if namespace != "" {
				meta.Namespaces = strings.Split(namespace, ",")
			}
			resources.Metadata = meta
		}

		report := audit.AuditRBACResources(resources)

		if jsonOut {
			jsonData, err := json.MarshalIndent(report, "", "  ")
			if err != nil {
				log.Fatalf("Failed to marshal audit report: %v", err)
			}
			filename := "rbac_audit_report.json"
			if err := os.WriteFile(filename, jsonData, 0644); err != nil {
				log.Fatalf("Failed to write audit report: %v", err)
			}
			fmt.Printf("Audit report written to %s\n", filename)
		} else {
			if len(report.Findings) == 0 {
				fmt.Println("No risky RBAC configurations found.")
				return
			}
			fmt.Println("Risky RBAC configurations:")
			for _, finding := range report.Findings {
				fmt.Printf("[%s] %s/%s", finding.Risk, finding.ResourceKind, finding.ResourceName)
				if finding.Namespace != "" {
					fmt.Printf(" (ns: %s)", finding.Namespace)
				}
				fmt.Printf(": %s\n", finding.Reason)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ruleAuditCmd)
	ruleAuditCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	ruleAuditCmd.Flags().StringVar(&namespace, "namespace", "", "Namespaces to audit (comma-separated)")
	ruleAuditCmd.Flags().BoolVar(&jsonOut, "jsonOut", false, "Output audit results to JSON file")
	ruleAuditCmd.Flags().StringVar(&inputFile, "input", "", "Path to a previously saved RBAC resources JSON file to audit")
}
