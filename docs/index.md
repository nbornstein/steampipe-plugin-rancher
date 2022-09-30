---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/rancher.svg"
brand_color: "#326CE5"
display_name: "Rancher"
short_name: "rancher"
description: "Steampipe plugin for Rancher components."
og_description: "Query Rancher with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/rancher-social-graphic.png"
---

# Rancher + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Rancher](https://rancher.io) is an open-source system for automating deployment, scaling, and management of containerized applications.

For example:

```sql
select
  name,
  namespace,
  phase,
  creation_timestamp,
  pod_ip
from
  rancher_pod;
```

```
+-----------------------------------------+-------------+-----------+---------------------+-----------+
| name                                    | namespace   | phase     | creation_timestamp  | pod_ip    |
+-----------------------------------------+-------------+-----------+---------------------+-----------+
| metrics-server-86cbb8457f-bf8dm         | kube-system | Running   | 2021-06-11 14:21:48 | 10.42.0.5 |
| coredns-7448499f4d-klb8l                | kube-system | Running   | 2021-06-11 14:21:48 | 10.42.0.6 |
| helm-install-traefik-crd-hb87d          | kube-system | Succeeded | 2021-06-11 14:21:48 | 10.42.0.3 |
| local-path-provisioner-5ff76fc89d-c9hnm | kube-system | Running   | 2021-06-11 14:21:48 | 10.42.0.2 |
+-----------------------------------------+-------------+-----------+---------------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/rancher/tables)**

## Get started

### Install

Download and install the latest Rancher plugin:

```bash
steampipe plugin install rancher
```

### Configuration

Installing the latest rancher plugin will create a config file (`~/.steampipe/config/rancher.spc`) with a single connection named `rancher`:

```hcl
connection "rancher" {
  plugin      = "rancher"
}
```

This will create a `rancher` connection that uses the default kubeconfig context.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-rancher
- Community: [Slack Channel](https://steampipe.io/community/join)

## Advanced configuration options

If you have a kube config setup using the kubectl CLI Steampipe just works with that connection.

The Rancher plugin allows you set the name of kube kubectl context with the `config_context` argument in any connection profile. You may also specify the path to kube config file with `config_path` argument.

This plugin also supports querying Rancher clusters using [OpenID Connect](https://rancher.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens) (OIDC) authentication. No extra configuration is required in a connection profile to query clusters using OIDC.

### Credentials via kube config

```hcl
connection "k8s_minikube" {
  plugin         = "rancher"
  config_context = "minikube"
  # config_path    = "~/.kube/config"
}
```
