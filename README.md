# Actions Documenter

This is a simple CLI tool that allows you to automatically generate **highly opinionated** documentation with 
minimal input from the human

## Installing

```shell
git clone git@github.com:userbradley/actions-documenter.git
go install
```

## Setting up the Configuration

Create a file called `readme.hcl`

In the file put the below:

```hcl
version = "v1.0.0"

title {
  enabled = true
  override = ""
}

quickstart {
  path = "docs/quickstart.md"
}

examples {
  example {
    enabled = true
    name    = "Example 1"
    path    = "docs/example1.md"
  }

}

footer {
  footer_from = "footer.md"
}
```

### How to use the version

To make life easier and not having to update every reference for the verison, you can use `@${{version}}` in the action file

```yaml
    steps:
      - uses: userbradley/example-action@${{version}}
```

When this gets rendered out, it will be replaced with what ever `version` is set to

You should ensure that the `readme.hcl` file is in the same location as the `action.yml` file

### How to add more examples

To add more examples, just repeat the `example` block like below

```hcl
examples {
  example {
    enabled = true
    name    = "Example 1"
    path    = "docs/example1.md"
  }
  example {
    enabled = true
    name    = "Example 2"
    path    = "docs/example2.md"
  }
}
```

## Contact

As per the license, if you want to use this internally please contact me on [legal@breadnet.co.uk](mailto:legal@breadnet.co.uk)
