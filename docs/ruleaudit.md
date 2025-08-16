# Rule Audit Command

The `ruleaudit` command audits Kubernetes RBAC resources for risky configurations using built-in rules. You can audit resources live from a cluster or from a previously saved JSON file.

## Usage

```sh
gorbac ruleaudit [flags]
```

**Flags:**

- `--kubeconfig`: Path to the kubeconfig file (optional).
- `--namespace`: Comma-separated list of namespaces to audit (optional).
- `--jsonOut`: Output the audit report to a JSON file.
- `--input`: Path to a previously saved RBAC resources JSON file to audit (optional).

**Examples:**

- Audit live cluster RBAC resources and output to JSON:

  ```sh
  ./gorbac ruleaudit --jsonOut
  ```

- Audit RBAC resources from a specific namespace:

  ```sh
  ./gorbac ruleaudit --namespace=my-namespace --jsonOut
  ```

- Audit from a previously saved RBAC resources file:

  ```sh
  ./gorbac ruleaudit --input=rbac_resources.json --jsonOut
  ```

## How It Works

- If `--input` is provided, the command reads RBAC resources from the specified JSON file.
- Otherwise, it fetches live RBAC resources from the cluster using the provided kubeconfig and namespaces.
- The command audits the resources for risky configurations (e.g., wildcard permissions, broad bindings) and outputs findings.
- If `--jsonOut` is set, the audit report is saved to `rbac_audit_report.json`.
- Otherwise, findings are printed to the console.

## Output

- **JSON Output:**
  - The audit report is saved as `rbac_audit_report.json`.
- **Console Output:**
  - A summary of risky RBAC configurations is printed. If no risky configurations are found, a message is displayed.

---

See the main README for more details on installation and usage.
