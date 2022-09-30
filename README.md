# Rancher Plugin for Steampipe

Use SQL to query Rancher components.

- **[Get started →](https://hub.steampipe.io/plugins/nbornstein/rancher)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/nbornstein/rancher/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/nbornstein/steampipe-plugin-rancher/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install rancher
```

Run a query:

```sql
select name, creation_timestamp, addresses, capacity from rancher_node;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/nbornstein/steampipe-plugin-rancher.git
cd steampipe-plugin-rancher
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/rancher.spc
```

Try it!

```
steampipe query
> .inspect rancher
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/nbornstein/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/nbornstein/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/nbornstein/steampipe-plugin-rancher/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/nbornstein/steampipe/labels/help%20wanted)
- [Rancher Plugin](https://github.com/nbornstein/steampipe-plugin-rancher/labels/help%20wanted)
