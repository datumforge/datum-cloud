# Contributing

Given external users will not have write to the branches in this repository, you'll need to follow the forking process to open a PR - [here](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork) is a guide from github on how to do so.

Please also read our main [contributing guide](https://github.com/datumforge/.github/blob/main/CONTRIBUTING.md) in addition to this one; the main guide mostly says that we'd like for you to open an issue first but it's not hard-required, and that we accept all forms of proposed changes given the state of this code base (in it's infancy, still!)

## Getting Started

Included in this repo is a [Taskfile](cmd/Taskfile.yaml) that makes the process quick. If you haven't used `task` before, head over to the upstream [docs](https://taskfile.dev/) to [install](https://taskfile.dev/installation/).

### Prerequisites

Before running any of the commands you will need to install the required dependencies
```bash
task install
```

### Pre-Commit Hooks

We have several `pre-commit` hooks that should be run before pushing a commit. Make sure this is installed:

```bash
brew install pre-commit
pre-commit install
```

You can optionally run against all files:

```bash
pre-commit run --all-files
```

## Datum Cloud Server

### Prerequisites

The `datum-cloud` server depends on the [datum-api](https://github.com/datumforge/datum) to be running because it makes requests to this server. In most cases, if you are running this server locally, you'll also want to run the `datum-api` locally as well. Refer to the [contributing guide](https://github.com/datumforge/datum/blob/main/.github/CONTRIBUTING.md#starting-the-server) in the Datum repo on getting started.

### Getting Started

Starting the server, you can use the cli to bring up the local server. This will create a token to authenticate to the datum-api and startup the server:

```bash
task run-dev
```

You can then use the cli to do things like, create a new workspace:

```bash
task cli:workspace:create
```

<details>
<summary>Workspace Creation</summary>

```
task cli:workspace:create

task: [cli:workspace:create] go run cmd/cli/main.go workspace create
Name: mitb
Description (optional): Description (optional): █
Domains (optional): datum.net
👉 Production & Testing
Environments:  [production testing]

> creating workspaces... 100% [===============]  [1s]
ID:  01J0M42YRA021QCM060505XD47
Name:  mitb
Domains:  datum.net
```
</details>

## Datum Cloud CLI

The `cli` included in this repo is used to talk to both the `datum-cloud` server and the `datum-api` server so in most cases you will need both servers running for the requests to be successful.
