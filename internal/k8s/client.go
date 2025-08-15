package k8s

// Client wraps Kubernetes client for RBAC operations
import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/flushthemoney/GoRBAC/internal/types"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Wrapper struct for k8s clientset
type Client struct {
	clientset *kubernetes.Clientset
	config    *rest.Config
}

// NewClient creates a new Kubernetes client
func NewClient(kubeconfig string) (*Client, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		// Try in-cluster config first
		config, err = rest.InClusterConfig()
		if err != nil {
			// Fall back to default kubeconfig location
			home := os.Getenv("HOME")
			if home == "" {
				home = os.Getenv("USERPROFILE")
			}
			kubeconfig = filepath.Join(home, ".kube/config")
			config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	return &Client{
		clientset: clientset,
		config:    config,
	}, nil
}

func (c *Client) GetRBACResources(ctx context.Context, namespace string) (*types.RBACResources, error) {
	resources := &types.RBACResources{}

	// Get Roles
	if namespace != "" {
		roles, err := c.getRoles(ctx, namespace)
		if err != nil {
			return nil, fmt.Errorf("failed to get roles: %w", err)
		}
		resources.Roles = roles
	}

	// Get ClusterRoles
	clusterRoles, err := c.getClusterRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster roles: %w", err)
	}
	resources.ClusterRoles = clusterRoles

	// Get RoleBindings
	if namespace != "" {
		roleBindings, err := c.getRoleBindings(ctx, namespace)
		if err != nil {
			return nil, fmt.Errorf("failed to get role bindings: %w", err)
		}
		resources.RoleBindings = roleBindings
	}

	// Get ClusterRoleBindings (always cluster-scoped)
	clusterRoleBindings, err := c.getClusterRoleBindings(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster role bindings: %w", err)
	}
	resources.ClusterRoleBindings = clusterRoleBindings

	return resources, nil

}

// getRoles retrieves roles from the cluster
func (c *Client) getRoles(ctx context.Context, namespace string) ([]rbacv1.Role, error) {
	if namespace != "" {
		roles, err := c.clientset.RbacV1().Roles(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		return roles.Items, nil
	}

	return []rbacv1.Role{}, nil
}

// getClusterRoles retrieves cluster roles from the cluster
func (c *Client) getClusterRoles(ctx context.Context) ([]rbacv1.ClusterRole, error) {
	clusterRoles, err := c.clientset.RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return clusterRoles.Items, nil
}

// getRoleBindings retrieves role bindings from the cluster
func (c *Client) getRoleBindings(ctx context.Context, namespace string) ([]rbacv1.RoleBinding, error) {
	if namespace != "" {
		roleBindings, err := c.clientset.RbacV1().RoleBindings(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		return roleBindings.Items, nil
	}

	return []rbacv1.RoleBinding{}, nil
}

// getClusterRoleBindings retrieves cluster role bindings from the cluster
func (c *Client) getClusterRoleBindings(ctx context.Context) ([]rbacv1.ClusterRoleBinding, error) {
	clusterRoleBindings, err := c.clientset.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return clusterRoleBindings.Items, nil
}
