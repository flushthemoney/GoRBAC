# RBACLens

RBACLens is a powerful CLI tool for auditing and analyzing Kubernetes RBAC (Role-Based Access Control) resources. It helps cluster administrators and security teams identify risky RBAC configurations, visualize permissions, and ensure best practices are followed.

It is also intended to be used by people trying to learn more about RBAC rules on K8s

![Demo](./docs/demo.gif)

## Features

- **Fetch RBAC resources**: Retrieve Roles, ClusterRoles, RoleBindings, and ClusterRoleBindings from a Kubernetes cluster and save them to a JSON file for offline analysis.
- **Audit for risky configurations**: Analyze RBAC resources using built-in rules to detect overly permissive or dangerous settings.
- **Flexible input**: Audit live from a cluster or from previously saved JSON files.

Comprehensive documentation for each command, including usage, flags, and examples, is available in the [docs/](docs/) directory.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

#### Option 1: Download Pre-built Binary (Recommended)

Pre-built binaries for Linux, macOS, and Windows are available on the [GitHub Releases page](https://github.com/flushthemoney/RBACLens/releases).

1. Go to the [releases page](https://github.com/flushthemoney/RBACLens/releases).
2. Download the appropriate binary for your OS and architecture.
3. Move the binary to a directory in your `PATH` (e.g., `~/bin`, `~/.local/bin`, or `/usr/local/bin`).
4. (Optional) Rename the binary to `rbaclens` for convenience. This makes it easier to run the tool from the terminal.

**Linux:**

```sh
mv RBACLens-linux-amd64 ~/bin/rbaclens
chmod +x ~/bin/rbaclens
```

**macOS:**

```sh
mv RBACLens-darwin-amd64 /usr/local/bin/rbaclens
chmod +x /usr/local/bin/rbaclens
```

#### Option 2: Install via Go

You can install RBACLens directly using Go (requires Go 1.16+):

```sh
go install github.com/flushthemoney/RBACLens@latest
```

This will place the binary (named `RBACLens` by default) in your `$GOPATH/bin`, `$GOBIN`, or `$HOME/go/bin` directory.

**Add Go bin directory to your PATH:**

- **Bash / Zsh:**

  ```sh
  export PATH="$PATH:$HOME/go/bin"
  ```

  Add the above line to your `~/.bashrc` or `~/.zshrc`.

- **Fish shell:**
  ```fish
  set -U fish_user_paths $fish_user_paths $HOME/go/bin
  ```

**Symlink for Convenience:**

If the binary is named `RBACLens`, you may want to symlink it to `rbaclens`:

```sh
ln -sf "$HOME/go/bin/RBACLens" "$HOME/go/bin/rbaclens"
```

After this, you can use the `rbaclens` command as described below.

#### Option 3: Build from Source

1. Clone the repository:
   ```sh
   git clone https://github.com/flushthemoney/RBACLens.git
   ```
2. Navigate to the project directory:
   ```sh
   cd RBACLens
   ```
3. Build the binary:
   ```sh
   go build -o rbaclens main.go
   ```

## Usage

## Commands Overview

- **Fetch RBAC Resources**: Use the `fetch` command to collect RBAC resources from your cluster. See [docs/fetch.md](docs/fetch.md) for details.
- **Audit RBAC Resources**: Use the `ruleaudit` command to analyze RBAC resources for risky configurations. See [docs/ruleaudit.md](docs/ruleaudit.md) for details.

For more information on all commands and advanced usage, refer to the [docs/](docs/) directory.
