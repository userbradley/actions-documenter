Mkdocs site builder and deployer

## Inputs

| Name | Description | Required | Default Value |
|------|-------------|----------|---------------|
| `serviceAccount` | Email address of the service account to use to Upload to the GCS Bucket | `true` | `Null` |
| `sensitive` | Should the Workload Identity provider use the Sensitive Pool | `false` | `false` |
| `GITHUB_TOKEN` | GitHub Token | `true` | `Null` |
| `siteUrl` | URL of the site to link the PR comment to | `true` | `Null` |
| `directory` | Location of the site to build | `true` | `Null` |
| `mkdocsVersion` | Version of MKdocs to install | `false` | `9.1.21` |
| `gcsBucket` | Name of the GCS Bucket to deploy the site to | `true` | `Null` |
