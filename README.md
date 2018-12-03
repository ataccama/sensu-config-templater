# Sensu Config Templater

> Allows you to share configured resources across Sensu namespaces

## Features

So far, it manages these resources:
- namespaces
- assets
- checks

We plan to add support for these very soon:
- filters
- handlers
- mutators
- groups

And some time later:
- hooks

### Namespace population

All resources align with the [reference model](https://docs.sensu.io/sensu-go/5.0/reference/) with one exception. It's the `namespace` parameter itself. This tool replaces it with `namespaces` (note the plural form) array argument. It's not mandatory - if you omit it, the resource will be populated to every configured namespace. Otherwise it allows you tailor the namespace(s) placement of the given resource without its duplication.

## Usage

1. Create the basic directory structure with `mkdir -p config/{checks,assets,namespaces}`
2. `cd config`
3. Create YAML configs for resources
4. `sensu-config-templater | sensuctl create`

## Building

Just run `go get github.com/ataccama/sensu-config-templater`
