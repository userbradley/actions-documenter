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
      - uses: userbradley/action-mkdocs@${{version}}
        with:
          gcsBucket: 
          serviceAccount: 
          directory: site
```

