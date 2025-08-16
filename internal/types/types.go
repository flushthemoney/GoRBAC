package types

import (
	"time"

	rbacv1 "k8s.io/api/rbac/v1"
)

// Metadata holds metadata about the fetch
type Metadata struct {
	ClusterName string    `json:"clusterName"`
	Timestamp   time.Time `json:"timestamp"`
	Namespaces  []string  `json:"namespaces,omitempty"`
}

// RBACResources holds all the RBAC resources.
type RBACResources struct {
	Metadata            Metadata                    `json:"metadata"`
	Roles               []rbacv1.Role               `json:"roles,omitempty"`
	ClusterRoles        []rbacv1.ClusterRole        `json:"clusterRoles,omitempty"`
	RoleBindings        []rbacv1.RoleBinding        `json:"roleBindings,omitempty"`
	ClusterRoleBindings []rbacv1.ClusterRoleBinding `json:"clusterRoleBindings,omitempty"`
}
