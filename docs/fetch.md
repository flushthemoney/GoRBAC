# Fetch Command

The `fetch` command retrieves Kubernetes RBAC resources for analysis or auditing. You can fetch resources live from a cluster and optionally save them to a JSON file.

## Usage

```sh
rbaclens fetch [flags]
```

**Flags:**

- `--kubeconfig`: Path to the kubeconfig file (optional).
- `--namespace`: Comma-separated list of namespaces to fetch (optional).
- `--json-out`: Output the RBAC resources to a JSON file.

**Examples:**

- Fetch all RBAC resources from the cluster and output to JSON:

  ```sh
  ./rbaclens fetch --json-out
  ```

- Fetch RBAC resources from a specific namespace:

  ```sh
  ./rbaclens fetch --namespace=my-namespace --json-out
  ```

- Fetch RBAC resources from multiple namespaces:

  ```sh
  ./rbaclens fetch --namespace=my-namespace,another-namespace --json-out
  ```

- Use a specific kubeconfig file:

  ```sh
  ./rbaclens fetch --kubeconfig=/path/to/kubeconfig --json-out
  ```

## How It Works

- Connects to the Kubernetes cluster using the provided kubeconfig (or default if not specified).
- Fetches RBAC resources from the specified namespaces (or all if not specified).
- If `--json-out` is set, the resources are saved to `rbac_resources.json`.
- Otherwise, resources are not saved to disk.

## Output

- **JSON Output:**
  - The RBAC resources are saved as `rbac_resources.json`.
- **Console Output:**
  - A success message is printed with the output file name.

---

See the main README for more details on installation and usage.
