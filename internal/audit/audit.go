package audit

import (
	"sort"

	"github.com/flushthemoney/GoRBAC/internal/types"
	v1 "k8s.io/api/rbac/v1"
)

type RiskLevel string

const (
	RiskLow    RiskLevel = "🔵 Low"
	RiskMedium RiskLevel = "🟡 Medium"
	RiskHigh   RiskLevel = "🔴 High"
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
				continue
			}
			if isRuleMediumRisk(rule) {
				findings = append(findings, AuditResult{
					ResourceKind: "ClusterRole",
					ResourceName: cr.Name,
					Risk:         RiskMedium,
					Reason:       mediumRiskReason(rule),
				})
				continue
			}
			if isRuleLowRisk(rule) {
				findings = append(findings, AuditResult{
					ResourceKind: "ClusterRole",
					ResourceName: cr.Name,
					Risk:         RiskLow,
					Reason:       lowRiskReason(rule),
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
				continue
			}
			if isRuleMediumRisk(rule) {
				findings = append(findings, AuditResult{
					ResourceKind: "Role",
					ResourceName: r.Name,
					Namespace:    r.Namespace,
					Risk:         RiskMedium,
					Reason:       mediumRiskReason(rule),
				})
				continue
			}
			if isRuleLowRisk(rule) {
				findings = append(findings, AuditResult{
					ResourceKind: "Role",
					ResourceName: r.Name,
					Namespace:    r.Namespace,
					Risk:         RiskLow,
					Reason:       lowRiskReason(rule),
				})
			}
		}
	}

	// Check for medium/low risk verbs in ClusterRoleBindings and RoleBindings (e.g., impersonate, escalate, bind)
	// (ClusterRoleBindings and RoleBindings themselves are high risk if bound to system groups, already checked above)

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

	// Sort findings by risk: High > Medium > Low
	sortFindingsByRisk(findings)
	return AuditReport{
		Metadata: resources.Metadata,
		Findings: findings,
	}
}

// sortFindingsByRisk sorts findings in-place by risk: High > Medium > Low
func sortFindingsByRisk(findings []AuditResult) {
	riskOrder := map[RiskLevel]int{
		RiskHigh:   0,
		RiskMedium: 1,
		RiskLow:    2,
	}
	sort.SliceStable(findings, func(i, j int) bool {
		return riskOrder[findings[i].Risk] < riskOrder[findings[j].Risk]
	})
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

// isRuleMediumRisk checks for common medium risk RBAC configurations
func isRuleMediumRisk(rule v1.PolicyRule) bool {
	// Granting get/list/watch on secrets
	for _, res := range rule.Resources {
		if res == "secrets" {
			for _, v := range rule.Verbs {
				if v == "get" || v == "list" || v == "watch" {
					return true
				}
			}
		}
	}
	// Granting create on workloads (pods, deployments, etc.)
	for _, res := range rule.Resources {
		if res == "pods" || res == "deployments" || res == "statefulsets" || res == "daemonsets" || res == "jobs" || res == "cronjobs" {
			for _, v := range rule.Verbs {
				if v == "create" {
					return true
				}
			}
		}
	}
	// Granting create on persistentvolumes
	for _, res := range rule.Resources {
		if res == "persistentvolumes" {
			for _, v := range rule.Verbs {
				if v == "create" {
					return true
				}
			}
		}
	}
	// Granting impersonate, escalate, or bind verbs
	for _, v := range rule.Verbs {
		if v == "impersonate" || v == "escalate" || v == "bind" {
			return true
		}
	}
	// Granting access to proxy subresource of nodes
	for _, res := range rule.Resources {
		if res == "nodes/proxy" {
			return true
		}
	}
	return false
}

// isRuleLowRisk checks for common low risk RBAC configurations
func isRuleLowRisk(rule v1.PolicyRule) bool {
	// Granting list/watch on non-sensitive resources cluster-wide
	for _, v := range rule.Verbs {
		if v == "list" || v == "watch" {
			for _, res := range rule.Resources {
				if res == "pods" || res == "services" || res == "configmaps" || res == "endpoints" {
					return true
				}
			}
		}
	}
	// Granting get on configmaps
	for _, res := range rule.Resources {
		if res == "configmaps" {
			for _, v := range rule.Verbs {
				if v == "get" {
					return true
				}
			}
		}
	}
	// Granting patch on namespaces
	for _, res := range rule.Resources {
		if res == "namespaces" {
			for _, v := range rule.Verbs {
				if v == "patch" {
					return true
				}
			}
		}
	}
	return false
}

// mediumRiskReason returns a string describing the medium risk found
func mediumRiskReason(rule v1.PolicyRule) string {
	for _, res := range rule.Resources {
		if res == "secrets" {
			for _, v := range rule.Verbs {
				if v == "get" || v == "list" || v == "watch" {
					return "Rule grants get/list/watch on secrets, which can leak sensitive data."
				}
			}
		}
		if res == "pods" || res == "deployments" || res == "statefulsets" || res == "daemonsets" || res == "jobs" || res == "cronjobs" {
			for _, v := range rule.Verbs {
				if v == "create" {
					return "Rule grants create on workloads (pods, deployments, etc.), which can lead to privilege escalation."
				}
			}
		}
		if res == "persistentvolumes" {
			for _, v := range rule.Verbs {
				if v == "create" {
					return "Rule grants create on persistentvolumes, which can allow hostPath abuse."
				}
			}
		}
		if res == "nodes/proxy" {
			return "Rule grants access to proxy subresource of nodes, which can allow bypassing audit and admission controls."
		}
	}
	for _, v := range rule.Verbs {
		if v == "impersonate" {
			return "Rule grants impersonate verb, which can allow privilege escalation."
		}
		if v == "escalate" {
			return "Rule grants escalate verb, which can allow privilege escalation."
		}
		if v == "bind" {
			return "Rule grants bind verb, which can allow privilege escalation."
		}
	}
	return "Rule is considered medium risk."
}

// lowRiskReason returns a string describing the low risk found
func lowRiskReason(rule v1.PolicyRule) string {
	for _, v := range rule.Verbs {
		if v == "list" || v == "watch" {
			for _, res := range rule.Resources {
				if res == "pods" || res == "services" || res == "configmaps" || res == "endpoints" {
					return "Rule grants list/watch on non-sensitive resources cluster-wide."
				}
			}
		}
	}
	for _, res := range rule.Resources {
		if res == "configmaps" {
			for _, v := range rule.Verbs {
				if v == "get" {
					return "Rule grants get on configmaps."
				}
			}
		}
		if res == "namespaces" {
			for _, v := range rule.Verbs {
				if v == "patch" {
					return "Rule grants patch on namespaces, which can affect pod security or network policies."
				}
			}
		}
	}
	return "Rule is considered low risk."
}
