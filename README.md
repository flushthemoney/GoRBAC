# GoRBAC

GoRBAC is a CLI tool for auditing Kubernetes RBAC resources.

## About

GoRBAC is a CLI tool designed to help you audit and analyze
Kubernetes RBAC (Role-Based Access Control) resources.

You can use GoRBAC to:

- Fetch RBAC resources from a cluster and save them to a JSON file for further analysis.
- Audit RBAC resources for risky configurations using built-in rules (see `ruleaudit` command).

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

The `fetch` command retrieves Roles, ClusterRoles, RoleBindings, and ClusterRoleBindings from a Kubernetes cluster and can save them to a JSON file for further analysis.

See [docs/fetch.md](docs/fetch.md) for full usage, flags, and examples.

---

### Audit RBAC Resources for Risky Configurations

The `ruleaudit` command audits RBAC resources for risky configurations using built-in rules. You can audit live from a cluster or from a previously saved JSON file.

See [docs/ruleaudit.md](docs/ruleaudit.md) for full usage, flags, and examples.
