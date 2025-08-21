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
var includeSystem bool

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

		report := audit.AuditRBACResourcesWithOptions(resources, audit.AuditOptions{
			IncludeSystemComponents: includeSystem,
		})

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
			printAuditReport(report)
		}
	},
}

func init() {
	rootCmd.AddCommand(ruleAuditCmd)
	ruleAuditCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	ruleAuditCmd.Flags().StringVar(&namespace, "namespace", "", "Namespaces to audit (comma-separated)")
	ruleAuditCmd.Flags().BoolVar(&jsonOut, "json-out", false, "Output audit results to JSON file")
	ruleAuditCmd.Flags().StringVar(&inputFile, "input", "", "Path to a previously saved RBAC resources JSON file to audit")
	ruleAuditCmd.Flags().BoolVar(&includeSystem, "include-system", false, "Include system components in audit results (may produce many findings)")
}

// printAuditReport prints a formatted audit report to the console
func printAuditReport(report audit.AuditReport) {
	fmt.Println("╭─────────────────────────────────────────────────────────────╮")
	fmt.Println("│                    🔍 RBAC Security Audit                   │")
	fmt.Println("╰─────────────────────────────────────────────────────────────╯")
	fmt.Println()

	// Print summary statistics
	fmt.Printf("📊 Audit Summary:\n")
	fmt.Printf("   • ClusterRoles:        %d\n", report.Summary.TotalClusterRoles)
	fmt.Printf("   • Roles:               %d\n", report.Summary.TotalRoles)
	fmt.Printf("   • ClusterRoleBindings: %d\n", report.Summary.TotalClusterRoleBindings)
	fmt.Printf("   • RoleBindings:        %d\n", report.Summary.TotalRoleBindings)
	fmt.Printf("   • System resources skipped: %d\n", report.Summary.SystemResourcesSkipped)
	fmt.Println()

	if report.Summary.TotalFindings == 0 {
		fmt.Println("✅ No security issues found!")
		fmt.Println("   All RBAC configurations appear to follow security best practices.")
		if report.Summary.SystemResourcesSkipped > 0 {
			fmt.Printf("   (Skipped %d system components - use --include-system to see them)\n", report.Summary.SystemResourcesSkipped)
		}
		return
	}

	// Print findings summary
	fmt.Printf("⚠️  Security Issues Found: %d\n", report.Summary.TotalFindings)
	if report.Summary.HighRiskFindings > 0 {
		fmt.Printf("   🔴 High Risk:   %d\n", report.Summary.HighRiskFindings)
	}
	if report.Summary.MediumRiskFindings > 0 {
		fmt.Printf("   🟡 Medium Risk: %d\n", report.Summary.MediumRiskFindings)
	}
	if report.Summary.LowRiskFindings > 0 {
		fmt.Printf("   🔵 Low Risk:    %d\n", report.Summary.LowRiskFindings)
	}
	fmt.Println()

	// Print detailed findings
	fmt.Println("📋 Detailed Findings:")
	fmt.Println("────────────────────────────────────────────────────────────────")

	for i, finding := range report.Findings {
		fmt.Printf("%d. [%s] %s/%s", i+1, finding.Risk, finding.ResourceKind, finding.ResourceName)
		if finding.Namespace != "" {
			fmt.Printf(" (namespace: %s)", finding.Namespace)
		}
		fmt.Println()
		fmt.Printf("   └─ %s\n", finding.Reason)
		if i < len(report.Findings)-1 {
			fmt.Println()
		}
	}

	fmt.Println()
	if report.Summary.SystemResourcesSkipped > 0 {
		fmt.Printf("💡 Tip: %d system resources were skipped. Use --include-system to include them.\n", report.Summary.SystemResourcesSkipped)
	}
}
