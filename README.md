# GoRBAC

GoRBAC is a CLI tool for auditing Kubernetes RBAC resources.

## About

GoRBAC is a CLI tool designed to help you audit and analyze
Kubernetes RBAC (Role-Based Access Control) resources. 

You can currently use GoRBAC to fetch RBAC resources from a cluster and 
save them to a JSON file for further analysis.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/flushthemoney/GoRBAC.git
   ```
2. Navigate to the project directory:
   ```sh
   cd GoRBAC
   ```
3. Build the binary:
   ```sh
   go build -o gorbac main.go
   ```

## Usage

### Fetch RBAC Resources

The `fetch` command fetches Roles, ClusterRoles, RoleBindings, and ClusterRoleBindings from a Kubernetes cluster.

```sh
gorbac fetch [flags]
```

**Flags:**

*   `--kubeconfig`: Path to the kubeconfig file (optional).
*   `--namespace`: Comma-separated list of namespaces to audit (optional).
*   `--jsonOut`: Output the RBAC resources to a JSON file.

**Examples:**

*   Fetch all RBAC resources from the cluster and output to a JSON file:

    ```sh
    ./gorbac fetch --jsonOut
    ```

*   Fetch RBAC resources from a specific namespace:

    ```sh
    ./gorbac fetch --namespace=my-namespace --jsonOut
    ```

*   Fetch RBAC resources from multiple namespaces:

    ```sh
    ./gorbac fetch --namespace=my-namespace,another-namespace --jsonOut
    ```

*   Use a specific kubeconfig file:

    ```sh
    ./gorbac fetch --kubeconfig=/path/to/kubeconfig --jsonOut
    ```