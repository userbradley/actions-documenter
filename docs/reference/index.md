---
title: readme.hcl reference
---

# `readme.hcl` reference

The configuration file is made up of several blocks
* Version
* Title
* Quickstart
* Examples
* Footer

## Version

This allows you to set the version of the module on release in the documentation

```terraform
version = "v1.0.0"
```

## Title

Allow you to omit the title or override it

```terraform
title {
  enabled = true
  override = ""
}
```

If you wish to override the title (So not pulling it from the action file) you can put it in the `override` section

## Quickstart

This section allows you to render a quickstart section at the top of the readme, allowing people to get on with their use very quickly.

This section is linking path to a file, based on the current directory of the command invocation

```terraform
quickstart {
  path = "docs/quickstart.md"
}
```

## Examples

This section allows you to include multiple examples.

The end result is rendered in the `README.md` so you can use any markdown or HTML that your source control system will render

```terraform
examples {
  example {
    enabled = true
    name    = "Example 1"
    path    = "docs/example1.md"
  }
}
```

??? tip "How to add multiple examples"
    Just repeat the `example` sections
    
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

## Footer

This section allows you to prepend stuff to the bottom of the file, like `Made By` or legal stuff

```terraform
footer {
  footer_from = "docs/footer.md"
}
```

## Full Reference

```terraform
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
  footer_from = "docs/footer.md"
}
```
