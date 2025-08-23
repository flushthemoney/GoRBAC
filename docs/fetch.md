# :inbox_tray: Fetch Command

!!! info
The `fetch` command retrieves Kubernetes RBAC resources for analysis or auditing. You can fetch resources live from a cluster and optionally save them to a JSON file for offline analysis.

---

## :hammer_and_wrench: Usage

```sh
rbaclens fetch [flags]
```

**Flags:**

- `--kubeconfig`: Path to the kubeconfig file (optional)
- `--namespace`: Comma-separated list of namespaces to fetch (optional)
- `--json-out`: Output the RBAC resources to a JSON file

---

## :bulb: Examples

- Fetch all RBAC resources from the cluster and output to JSON:

  ```
  rbaclens fetch --json-out
  ```

- Fetch RBAC resources from a specific namespace:

  ```
  rbaclens fetch --namespace=my-namespace --json-out
  ```

- Fetch RBAC resources from multiple namespaces:

  ```
  rbaclens fetch --namespace=my-namespace,another-namespace --json-out
  ```

- Use a specific kubeconfig file:

  ```
  rbaclens fetch --kubeconfig=/path/to/kubeconfig --json-out
  ```

---

## :gear: How It Works

1. Connects to the Kubernetes cluster using the provided kubeconfig (or default if not specified).
2. Fetches RBAC resources from the specified namespaces (or all if not specified).
3. If `--json-out` is set, the resources are saved to `rbac_resources.json`.
4. Otherwise, resources are not saved to disk.

---

## :package: Output

- **JSON Output:** The RBAC resources are saved as `rbac_resources.json`.
- **Console Output:** A success message is printed with the output file name.

---

!!! note
See the main [README](https://github.com/flushthemoney/RBACLens#readme) for more details on installation and usage.
