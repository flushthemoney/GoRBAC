# Rule Audit Command

The `ruleaudit` command audits Kubernetes RBAC resources for risky configurations using built-in rules. You can audit resources live from a cluster or from a previously saved JSON file.

This is designed to focus on **actual security issues** by filtering out legitimate system components, making the results actually meaningful.

## Usage

```sh
rbaclens ruleaudit [flags]
```

**Flags:**

- `--kubeconfig`: Path to the kubeconfig file (optional).
- `--namespace`: Comma-separated list of namespaces to audit (optional).
- `--json-out`: Output the audit report to a JSON file (optional).
- `--input`: Path to a previously saved RBAC resources JSON file to audit (optional).
- `--include-system`: Include system components in audit results (may produce many findings, disabled by default).

**Examples:**

- Audit live cluster RBAC resources with clean output:

  ```sh
  ./rbaclens ruleaudit
  ```

- Audit live cluster RBAC resources and output to JSON:

  ```sh
  ./rbaclens ruleaudit --json-out
  ```

- Audit RBAC resources from a specific namespace:

  ```sh
  ./rbaclens ruleaudit --namespace=my-namespace
  ```

- Audit from a previously saved RBAC resources file:

  ```sh
  ./rbaclens ruleaudit --input=rbac_resources.json
  ```

- Include system components in the audit (comprehensive scan):

  ```sh
  ./rbaclens ruleaudit --include-system
  ```

## Smart Filtering

By default, the audit tool **filters out system components** to focus on user-created or potentially problematic configurations:

### What Gets Filtered:

- **System Roles**: All roles prefixed with `system:` (e.g., `system:controller:*`, `system:kube-*`)
- **Default Kubernetes Roles**: `cluster-admin`, `admin`, `edit`, `view`
- **System Namespaces**: `kube-system`, `kube-public`, `kube-node-lease`, `default`
- **Legitimate System Bindings**: Known-good bindings like `system:public-info-viewer`
- **Service Account Bindings**: Legitimate system service account bindings

### What Gets Detected:

- **Custom roles** with overly broad permissions (`*` verbs or resources)
- **Dangerous bindings** to `system:unauthenticated` users
- **Risky custom permissions** on secrets, workloads, or persistent volumes
- **Custom roles** with privilege escalation potential

## Risk Levels

The audit categorizes findings into three risk levels:

- 🔴 **High Risk**: Wildcard permissions (`*`), bindings to unauthenticated users
- 🟡 **Medium Risk**: Access to secrets, workload creation, privilege escalation verbs
- 🔵 **Low Risk**: Broad list/watch permissions, configuration access

## How It Works

1. **Resource Collection:**

   - If `--input` is provided, reads RBAC resources from the specified JSON file
   - Otherwise, fetches live RBAC resources from the cluster using kubeconfig

2. **Smart Analysis:**

   - Applies intelligent filtering to focus on user-created resources
   - Evaluates each rule against security best practices
   - Categorizes findings by risk level

3. **Output Generation:**
   - Provides comprehensive statistics about the audit
   - Shows detailed findings with explanations
   - Offers clean, actionable results

## Output Formats

### CLI Output (Default)

The console output provides a clean, formatted report of the findings:

```
╭─────────────────────────────────────────────────────────────╮
│                    🔍 RBAC Security Audit                   │
╰─────────────────────────────────────────────────────────────╯

📊 Audit Summary:
   • ClusterRoles:        68
   • Roles:               12
   • ClusterRoleBindings: 56
   • RoleBindings:        12
   • System resources skipped: 146

✅ No security issues found!
   All RBAC configurations appear to follow security best practices.
   (Skipped 146 system components - use --include-system to see them)
```

If issues are found:

```
⚠️  Security Issues Found: 3
   🔴 High Risk:   1
   🟡 Medium Risk: 2

📋 Detailed Findings:
────────────────────────────────────────────────────────────────
1. [🔴 High] ClusterRole/dangerous-role
   └─ ClusterRole grants '*' verbs or resources, which is highly privileged.

2. [🟡 Medium] ClusterRole/secrets-reader
   └─ Rule grants get/list/watch on secrets, which can leak sensitive data.
```

### JSON Output

When using `--json-out`, the audit report is saved to `rbac_audit_report.json` with comprehensive metadata:

```json
{
  "metadata": {
    "clusterName": "my-cluster",
    "timestamp": "2025-08-21T12:00:00Z"
  },
  "findings": [...],
  "summary": {
    "totalClusterRoles": 68,
    "totalRoles": 12,
    "totalClusterRoleBindings": 56,
    "totalRoleBindings": 12,
    "totalFindings": 0,
    "highRiskFindings": 0,
    "mediumRiskFindings": 0,
    "lowRiskFindings": 0,
    "systemResourcesSkipped": 146
  }
}
```

## Best Practices

1. **Regular Auditing**: Run the audit regularly to catch configuration drift
2. **Start Simple**: Use the default filtering to focus on actionable issues
3. **Comprehensive Review**: Use `--include-system` occasionally for full cluster assessment
4. **Track Changes**: Save audit reports over time to track security improvements
5. **Focus on High/Medium**: Prioritize fixing high and medium risk findings first

---

See the main README for more details on installation and usage.
