[![Go Report Card](https://goreportcard.com/badge/github.com/datumforge/datum-cloud)](https://goreportcard.com/report/github.com/datumforge/datum-cloud)
[![Build status](https://badge.buildkite.com/9d99bb1f92d9195776d9983bea1f74314fd912706244c48863.svg)](https://buildkite.com/datum/datum-cloud)
[![Go Reference](https://pkg.go.dev/badge/github.com/datumforge/datum-cloud.svg)](https://pkg.go.dev/github.com/datumforge/datum-cloud)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache2.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)

# Datum Cloud

Building a SaaS offering on top of [Datum](https://github.com/datumforge/datum)

## Datum Cloud Server

The Datum Cloud server is used to consume the [Datum Server](https://github.com/datumforge/datum/) and apply an opinionated implementation on top of the the generics provided. Many, if not all, of the endpoints provided by the server use the [Datum Client](https://github.com/datumforge/datum/blob/main/pkg/datumclient/client.go#L21) to make requests to the Datum Server.

As an example, the `v1/workspaces` endpoint uses the `Datum` client to create an organizational hierarchy, called a `Workspace`:

```
│   └── rootorg <--- top level organization
│       ├── production <-- top level environment per customer organization
│       │   ├── assets <-- buckets
│       │   ├── customers
│       │   ├── orders
│       │   ├── relationships
│       │   │   ├── internal_users  <-- relationships
│       │   │   ├── marketing_subscribers
│       │   │   ├── marketplaces
│       │   │   ├── partners
│       │   │   └── vendors
│       │   └── sales
│       └── test <-- organization identical to production just named
```

## Datum Cloud CLI

The Datum Cloud cli is used to interact with the Datum Cloud Server as well as some requests directly to the Datum Server using the [Datum Client](https://github.com/datumforge/datum/blob/main/pkg/datumclient/client.go#L21). In order to use the cli, you must have a registered user with the Datum Server.

### Installation

```bash
brew install datumforge/tap/datum-cloud
```

### Upgrade

```bash
brew upgrade datumforge/tap/datum-cloud
```

### Usage

```bash
datum-cloud
the datum-cloud cli

Usage:
  datum-cloud [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  seed        the subcommands for creating demo data in datum
  workspace   the subcommands for working with the datum workspace
```

## Seeding Data

The `datum-cloud` cli has functionality to generate and load test data into `datum` using the `seed` command.

```bash
Usage:
  datum-cloud seed [command]

Available Commands:
  generate    generate random data for seeded environment
  init        init a new datum seeded environment
```

#### Using the Taskfile

On a brand new database, you should run:

1. Create a new user to authenticate with the Datum API, this command will fail on subsequent tries because the user will already exist.
    ```bash
    task register
    ```
1. Login as the user, create a new Personal Access Token that will be used to seed the data, generate a new data set, bulk load objects into the Datum API:
    ```bash
    task cli:seed:all
    ```

If instead, you prefer to use the CLI commands directly, keep reading.

### Generate Data

Using the `generate` subcommand, new random data will be stored in csv files:

```bash
datum-cloud seed generate
```

<details>
<summary>Generated Data</summary>

```bash
tree demodata
demodata
├── groups.csv
├── invites.csv
├── orgs.csv
└── users.csv
```

</details>

### Init Environment

Using the `init` subcommand, the data in the specified directory (defaults to `demodata` in the current directory), the csv files will be used to generate the data.

```bash
datum-cloud seed init
```

The newly created objects will be displayed when complete:

<details>
<summary>Results</summary>

```bash
> seeded environment created 100% [===============]  [3s]
Seeded Environment Created:
+--------------------------------------------------------------------------------------+
| Organization                                                                         |
+----------------------------+--------+-------------+-------------+----------+---------+
| ID                         | NAME   | DESCRIPTION | PERSONALORG | CHILDREN | MEMBERS |
+----------------------------+--------+-------------+-------------+----------+---------+
| 01J06RPZ8HQRWW4AZERHKWT2YH | Plus-U |             | false       |        0 |       1 |
+----------------------------+--------+-------------+-------------+----------+---------+
...
```

</details>

## Contributing

Please read the [contributing](.github/CONTRIBUTING.md) guide as well as the [Developer Certificate of Origin](https://developercertificate.org/). You will be required to sign all commits to the Datum project, so if you're unfamiliar with how to set that up, see [github's documentation](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification).

## Licensing

This repository contains `datum-cloud` which is open source software under [Apache 2.0](LICENSE). Datum-Cloud is a product produced from this open source software exclusively by Datum Technology, Inc. This product is produced under our published commercial terms (which are subject to change), and any logos or trademarks in this repository or the broader [datumforge](https://github.com/datumforge) organization are not covered under the Apache License.

Others are allowed to make their own distribution of this software or include this software in other commercial offerings, but cannot use any of the Datum logos, trademarks, cloud services, etc.

## Security

We take the security of our software products and services seriously, including all of the open source code repositories managed through our Github Organizations, such as [datumforge](https://github.com/datumforge). If you believe you have found a security vulnerability in any of our repositories, please report it to us through coordinated disclosure.

**Please do NOT report security vulnerabilities through public github issues, discussions, or pull requests!**

Instead, please send an email to `security@datum.net` with as much information as possible to best help us understand and resolve the issues. See the security policy attached to this repository for more details.

## Questions?

You can email us at `info@datum.net`, open a github issue in this repository, or reach out to [matoszz](https://github.com/matoszz) directly.



