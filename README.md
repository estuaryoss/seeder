Used to deploy test environments on [Estuary-Deployer(s)](https://github.com/estuaryoss/estuary-deployer)

Inspired by [Terraform](https://www.terraform.io/docs/cli/commands/index.html) and Kubernetes.

## Prerequisites
-  You deployed a net of Deployers

## Scope

- You want to seed your deployments across Deployer(s)

## Config precedence

The Deployer(s) services can be referred with one of these methods:

- get deployer apps from eureka server(s)
- get deployer apps from the discovery(ies)
- specify them one by one in config

The precedence is (from weaker to stronger):

`eureka < discovery < deployer`

Example:

```yaml
deploy_policy: fill #fill/robin
access_token: None
#eureka:
#  - "http://192.168.0.11:8080/eureka/v2"
#  - "http://192.168.0.11:8081/eureka/v2"
discovery:
  - "http://192.168.0.11:8082/"
  - "http://192.168.0.11:8083/"
#deployer:
#  - "http://192.168.0.11:8084/docker/"
#  - "http://192.168.0.11:8085/docker/"
```

## Katacoda

[seeder on Katacoda](https://www.katacoda.com/estuaryoss)

## Seeder CLI

Examples:

```bash
> seeder --help
Usage: seeder [--version] [--help] <command> [<args>]

Available commands are:
    apply       Usage: seeder apply
    destroy     Usage: seeder destroy
    init        Usage: seeder init
    plan        Usage: seeder plan
    show        Usage: seeder show
    validate    Usage: seeder validate
    version     Usage: seeder version


```

