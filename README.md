# GitHub Action: Example Action

A simple example action to demonstrate the output from the CLI

## Quickstart

```yaml
on: [push]
name: Build

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    permissions:
      id-token: write
      contents: read
    name: Build and Deploy
    steps:
      - uses: userbradley/example-action@v1.0.0
        with:
          gcsBucket:
          serviceAccount:
          directory: site
```
## Inputs

| Name | Description | Required | Default Value |
|------|-------------|----------|---------------|
| `mkdocsVersion` | Version of MKdocs to install | `false` | `9.1.21` |
| `gcsBucket` | Name of the GCS Bucket to deploy the site to | `true` | `Null` |
| `serviceAccount` | Email address of the service account to use to Upload to the GCS Bucket | `true` | `Null` |
| `sensitive` | Should the Workload Identity provider use the Sensitive Pool | `false` | `false` |
| `GITHUB_TOKEN` | GitHub Token | `true` | `Null` |
| `siteUrl` | URL of the site to link the PR comment to | `true` | `Null` |
| `directory` | Location of the site to build | `true` | `Null` |

## Examples

### Example 1

```yaml
on: [push]
name: Build

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    permissions:
      id-token: write
      contents: read
    name: Build and Deploy
    steps:
      - uses: userbradley/example-action@v1.0.0
        with:
          gcsBucket: 
          serviceAccount: 
          directory: site
```

---
This code is internal and not for release