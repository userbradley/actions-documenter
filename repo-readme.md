# GitHub Actions Automatic Documenter

## What problem this solves

An itch at the back of my brain

## How to use

You will need to install the CLI locally (Coming soon)

For the time being this can be done by running:

```shell
git clone git@github.com:userbradley/actions-documenter.git
go install
```

Once installed run the below

```shell
touch readme.hcl
mkdir examples
touch examples/basic.md
touch examples/quickstart.md
```

Ensure you run the command line tool from the same directory as the `actions.yml` file

## Generate the README

```shell
actions-documenter 
```

This will create the README file