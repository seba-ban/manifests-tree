# Manifests Tree CLI

CLI for displaying info about k8s resources in given yaml files. Work in progress.

```bash
manifests-tree --include-kinds=depl https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml 
```
```
.
└── apps/v1
    └── deployment [3]
        ├── cert-manager
        │   └── https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml:5384
        ├── cert-manager-cainjector
        │   └── https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml:5330
        └── cert-manager-webhook
            └── https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml:5452
```

## Output Formats

For available outputs, check out help for the `-o, --output` flag.

## Possible Inputs

There can be one or more inputs which can be:

* a path to a yaml file, e.g.:

    ```bash
    manifests-tree my-manifests.yaml
    ```

* path to a directory containing yaml files, e.g.:

    ```bash
    manifests-tree manifests_dir/
    ```

    In such case, directory will be searched for `.yaml` and `.yml` files. If `recursive` option is set to `true`, walk will enter all directories recursively.

* http(s) urls, e.g.:

    ```bash
    manifests-tree https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml
    ```

* stdin, e.g.:

    ```bash
    kustomize build . | manifests-tree -
    ```

    ```bash
    helm template \
        cert-manager cert-manager \
        --repo https://charts.jetstack.io \
        --namespace cert-manager \
        --create-namespace \
        --version v1.13.3 \
        | manifests-tree -
    ```

## Filtering

Filtering can be done on resource names, api versions, and kinds. For now only strict matching, and substring matching are supoported. Cf. `manifests-tree --help`.
