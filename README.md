# go-examples

A repository to play around with go, create examples, and collate useful snippets found online.

## Requirements

| Tool                                                                                     | Description                                     |
|:-----------------------------------------------------------------------------------------|:------------------------------------------------|
| [pre-commit](https://pre-commit.com/)                                                    | Used to ensure standards prior to commits       |

## Setup

### pre-commit

[pre-commit](https://pre-commit.com/) is used to enforce standards on this repository prior to committing any changes. This forms part of
our [Contributing](../CONTRIBUTING.md) standards. Please also see the
[pre-commit-config.yaml](../.pre-commit-config.yaml) file.

This is installed via the Pipfile, though this has to be initialised within this repository by running the below
command:

```shell
pre-commit install
```

# References

Below are a list of notable references in creating and structuring the examples
in this repository:

* [Alex Edwards - An introduction to Packages, Imports and Modules in Go](https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules#packages)
* [Uzair Ahmed - Simple Algorithm and Data Structure in Golang](https://medium.com/@uzairahmed01/algorithm-and-data-structure-in-golang-2869da82723e)
