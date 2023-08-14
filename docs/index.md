---
title: GitHub actions file Documenter
---

# GitHub actions file Documenter

## What is this

This is a Command line tool to automatically document (With minimal input) GitHub actions workflow files

## Quick start

### Install

```shell
go install github.com/userbradley/actions-documenter@latest
```

### Create Configuration file

```shell
curl -o readme.hcl https://raw.githubusercontent.com/userbradley/actions-documenter/main/example/readme.hcl
```

### Create quickstart file

```shell
mkdir docs
touch docs/quickstart.md
```
In this file you should put the minimum viable product to get it working.

!!! tip "Dynamically render the version"
    You can replace the version (so `@v1.0.0`) with `${{version}}` so when the command line tool is run, it will 
    find and replace so you don't need to update it everywhere when releasing a new version

