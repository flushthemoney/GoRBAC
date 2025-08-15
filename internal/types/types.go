package types

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

type RBACResources struct {
	Roles               []rbacv1.Role               `json:"roles,omitempty"`
	ClusterRoles        []rbacv1.ClusterRole        `json:"clusterRoles,omitempty"`
	RoleBindings        []rbacv1.RoleBinding        `json:"roleBindings,omitempty"`
	ClusterRoleBindings []rbacv1.ClusterRoleBinding `json:"clusterRoleBindings,omitempty"`
}
