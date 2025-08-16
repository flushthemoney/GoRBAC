package audit

import (
	"github.com/flushthemoney/GoRBAC/internal/types"
	v1 "k8s.io/api/rbac/v1"
)

type RiskLevel string

const (
	RiskLow    RiskLevel = "🔵Low"
	RiskMedium RiskLevel = "🟡Medium"
	RiskHigh   RiskLevel = "🔴High"
)

type AuditResult struct {
	ResourceKind string    `json:"resourceKind"`
	ResourceName string    `json:"resourceName"`
	Namespace    string    `json:"namespace,omitempty"`
	Risk         RiskLevel `json:"risk"`
	Reason       string    `json:"reason"`
}

type AuditReport struct {
	Metadata types.Metadata `json:"metadata"`
	Findings []AuditResult  `json:"findings"`
}

// AuditRBACResources audits the RBAC resources for risky configurations
func AuditRBACResources(resources types.RBACResources) AuditReport {
	findings := []AuditResult{}

	// Check ClusterRoles for risky rules
	for _, cr := range resources.ClusterRoles {
		for _, rule := range cr.Rules {
			if isRuleRisky(rule) {
				findings = append(findings, AuditResult{
					ResourceKind: "ClusterRole",
					ResourceName: cr.Name,
					Risk:         RiskHigh,
					Reason:       "ClusterRole grants '*' verbs or resources, which is highly privileged.",
				})
			}
		}
	}

	// Check Roles for risky rules
	for _, r := range resources.Roles {
		for _, rule := range r.Rules {
			if isRuleRisky(rule) {
				findings = append(findings, AuditResult{
					ResourceKind: "Role",
					ResourceName: r.Name,
					Namespace:    r.Namespace,
					Risk:         RiskHigh,
					Reason:       "Role grants '*' verbs or resources, which is highly privileged.",
				})
			}
		}
	}

	// Check ClusterRoleBindings for binding to system:serviceaccounts or system:unauthenticated
	for _, crb := range resources.ClusterRoleBindings {
		for _, s := range crb.Subjects {
			if s.Kind == "Group" && (s.Name == "system:unauthenticated" || s.Name == "system:serviceaccounts") {
				findings = append(findings, AuditResult{
					ResourceKind: "ClusterRoleBinding",
					ResourceName: crb.Name,
					Risk:         RiskHigh,
					Reason:       "ClusterRoleBinding grants cluster-wide access to all service accounts or unauthenticated users.",
				})
			}
		}
	}

	// Check RoleBindings for binding to system:unauthenticated
	for _, rb := range resources.RoleBindings {
		for _, s := range rb.Subjects {
			if s.Kind == "Group" && s.Name == "system:unauthenticated" {
				findings = append(findings, AuditResult{
					ResourceKind: "RoleBinding",
					ResourceName: rb.Name,
					Namespace:    rb.Namespace,
					Risk:         RiskHigh,
					Reason:       "RoleBinding grants access to unauthenticated users.",
				})
			}
		}
	}

	return AuditReport{
		Metadata: resources.Metadata,
		Findings: findings,
	}
}

// isRuleRisky checks if a policy rule is risky (e.g., grants '*')
func isRuleRisky(rule v1.PolicyRule) bool {
	for _, v := range rule.Verbs {
		if v == "*" {
			return true
		}
	}
	for _, r := range rule.Resources {
		if r == "*" {
			return true
		}
	}
	return false
}
