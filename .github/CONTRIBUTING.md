# Contributing

Given external users will not have write to the branches in this repository, you'll need to follow the forking process to open a PR - [here](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork) is a guide from github on how to do so.

Please also read our main [contributing guide](https://github.com/datumforge/.github/blob/main/CONTRIBUTING.md) in addition to this one; the main guide mostly says that we'd like for you to open an issue first but it's not hard-required, and that we accept all forms of proposed changes given the state of this code base (in it's infancy, still!)

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
