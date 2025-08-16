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

var ruleKubeconfig string
var ruleNamespaces string
var ruleJsonOut bool
var ruleInputFile string

// ruleauditCmd represents the ruleaudit command

var ruleauditCmd = &cobra.Command{
	Use:   "ruleaudit",
	Short: "Audit RBAC resources for risky configurations",
	Long: `Audit RBAC resources for risky configurations using built-in rules.
You can fetch live from a cluster or audit a previously saved JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var resources types.RBACResources
		if ruleInputFile != "" {
			data, err := os.ReadFile(ruleInputFile)
			if err != nil {
				log.Fatalf("Failed to read input file: %v", err)
			}
			if err := json.Unmarshal(data, &resources); err != nil {
				log.Fatalf("Failed to unmarshal input file: %v", err)
			}
		} else {
			clientset, err := k8s.NewClient(ruleKubeconfig)
			if err != nil {
				log.Fatalf("Failed to create Kubernetes client: %v", err)
			}
			rsrcPtr, err := clientset.GetRBACResources(cmd.Context(), ruleNamespaces)
			if err != nil {
				log.Fatalf("Failed to get RBAC resources: %v", err)
			}
			resources = *rsrcPtr
			meta := types.Metadata{
				ClusterName: "my-cluster", // Placeholder
				Timestamp:   time.Now(),
			}
			if ruleNamespaces != "" {
				meta.Namespaces = strings.Split(ruleNamespaces, ",")
			}
			resources.Metadata = meta
		}

		report := audit.AuditRBACResources(resources)

		if ruleJsonOut {
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
	rootCmd.AddCommand(ruleauditCmd)
	ruleauditCmd.Flags().StringVar(&ruleKubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	ruleauditCmd.Flags().StringVar(&ruleNamespaces, "namespace", "", "Namespaces to audit (comma-separated)")
	ruleauditCmd.Flags().BoolVar(&ruleJsonOut, "jsonOut", false, "Output audit results to JSON file")
	ruleauditCmd.Flags().StringVar(&ruleInputFile, "input", "", "Path to a previously saved RBAC resources JSON file to audit")
}
