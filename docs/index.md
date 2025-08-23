# RBACLens

<p align="center">
    <img src="./demo.gif" alt="RBACLens Demo"/>
</p>

---

RBACLens is a powerful CLI tool for auditing and analyzing Kubernetes RBAC (Role-Based Access Control) resources. It helps cluster administrators and security teams identify risky RBAC configurations, visualize permissions, and ensure best practices are followed.

!!! note
    RBACLens is also great for anyone learning about Kubernetes RBAC rules!

## :rocket: Features

- :inbox_tray: **Fetch RBAC resources**: Retrieve Roles, ClusterRoles, RoleBindings, and ClusterRoleBindings from a Kubernetes cluster and save them to a JSON file for offline analysis.
- :shield: **Audit for risky configurations**: Analyze RBAC resources using built-in rules to detect overly permissive or dangerous settings.
- :arrows_counterclockwise: **Flexible input**: Audit live from a cluster or from previously saved JSON files.

---

## :package: Installation

=== "Pre-built Binary (Recommended)"

    1. Go to the [GitHub Releases page](https://github.com/flushthemoney/RBACLens/releases)
    2. Download the appropriate binary for your OS and architecture.
    3. Move the binary to a directory in your `PATH` (e.g., `~/bin`, `~/.local/bin`, or `/usr/local/bin`).
    4. (Optional) Rename the binary to `rbaclens` for convenience. This makes it easier to run the tool from the terminal.

    === "Linux"

        ```sh
        mv RBACLens-linux-amd64 ~/bin/rbaclens
        chmod +x ~/bin/rbaclens
        ```

    === "macOS"

        ```sh
        mv RBACLens-darwin-amd64 /usr/local/bin/rbaclens
        chmod +x /usr/local/bin/rbaclens
        ```

    === "Windows (PowerShell)"

        ```powershell
        Rename-Item -Path .\RBACLens-windows-amd64.exe -NewName rbaclens.exe
        Move-Item -Path .\rbaclens.exe -Destination $env:USERPROFILE\bin\rbaclens.exe
        ```

    !!! tip
        You can rename the binary to `rbaclens` (lowercase) for convenience, and ensure the directory is in your `PATH` so you can run `rbaclens` from anywhere.

=== "Install via Go"

    ```sh
    go install github.com/flushthemoney/RBACLens@latest
    ```

    Ensure `$HOME/go/bin` is in your `PATH`.

    === "Fish shell"

        ```fish
        set -U fish_user_paths $fish_user_paths $HOME/go/bin
        ```

    === "Bash / Zsh"

        Add this to your `~/.bashrc` or `~/.zshrc`:

        ```sh
        export PATH="$PATH:$HOME/go/bin"
        ```

    === "Symlink for Convenience"

        If the binary is named `RBACLens`, you may want to symlink it to `rbaclens`:

        ```sh
        ln -sf "$HOME/go/bin/RBACLens" "$HOME/go/bin/rbaclens"
        ```

=== "Build from Source"

    ```sh
    git clone https://github.com/flushthemoney/RBACLens.git
    cd RBACLens
    go build -o rbaclens main.go
    ```

---

## :hammer_and_wrench: Usage

RBACLens provides the following commands:

- **Fetch RBAC Resources**: `rbaclens fetch`  
  [See details →](fetch.md)
- **Audit RBAC Resources**: `rbaclens ruleaudit`  
  [See details →](ruleaudit.md)

For advanced usage and all options, see the [project README](https://github.com/flushthemoney/RBACLens#readme).

---

## :books: Documentation

- [Fetch Command](fetch.md)
- [Rule Audit Command](ruleaudit.md)
- [Project README](https://github.com/flushthemoney/RBACLens#readme)

---

!!! info
    You can contribute or report issues on [GitHub](https://github.com/flushthemoney/RBACLens)
