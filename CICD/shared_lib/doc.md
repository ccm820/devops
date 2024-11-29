```plaintext
repo/
├── infra/
│   └── helm/
│       ├── Chart.yaml
│       ├── templates/
│       ├── values.yaml
│       └── environments/
│           ├── dev/
│           │   └── values.yaml
│           ├── sit/
│           │   └── values.yaml
│           └── prod/
│               └── values.yaml
```



Chart.yaml: Metadata for the Helm chart.
templates/: Kubernetes resource templates.
values.yaml: Global default values.
environments/: Environment-specific configuration files (dev, sit, prod).