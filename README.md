# GoRBAC

GoRBAC is a powerful CLI tool for auditing and analyzing Kubernetes RBAC (Role-Based Access Control) resources. It helps cluster administrators and security teams identify risky RBAC configurations, visualize permissions, and ensure best practices are followed.

## Features

- **Fetch RBAC resources**: Retrieve Roles, ClusterRoles, RoleBindings, and ClusterRoleBindings from a Kubernetes cluster and save them to a JSON file for offline analysis.
- **Audit for risky configurations**: Analyze RBAC resources using built-in rules to detect overly permissive or dangerous settings.
- **Flexible input**: Audit live from a cluster or from previously saved JSON files.

Comprehensive documentation for each command, including usage, flags, and examples, is available in the [docs/](docs/) directory.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

#### Option 1: Install via Go

You can install GoRBAC directly using Go (requires Go 1.16+):

```sh
go install github.com/flushthemoney/GoRBAC@latest
```

This will place the `gorbac` binary in your `$GOPATH/bin` or `$GOBIN` directory.

#### Option 2: Build from Source

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

## Commands Overview

- **Fetch RBAC Resources**: Use the `fetch` command to collect RBAC resources from your cluster. See [docs/fetch.md](docs/fetch.md) for details.
- **Audit RBAC Resources**: Use the `ruleaudit` command to analyze RBAC resources for risky configurations. See [docs/ruleaudit.md](docs/ruleaudit.md) for details.

For more information on all commands and advanced usage, refer to the [docs/](docs/) directory.
