# takolabel

`takolabel` is a CLI tool for manipulating GitHub Labels across multiple repositories.

## Installation

### macOS

```console
$ brew install tnagatomi/tap/takolabel
```

### Other platforms

Download packaged binaries from [Releases page](https://github.com/tnagatomi/takolabel/releases) in this repository.

## Usage

Set environment variables below (via `export` command etc.).

- TAKOLABEL_TOKEN (required)
  - GitHub Personal access token. A token with `repo` scope is needed if it will be run on a private repository. `public_repo` scope will suffice if it's a public repository.
- TAKOLABEL_HOST (optional)
  - Set this variable to the host URL (e.g. `https://ghe.example.com/`) if you want to work with repositories hosted on GitHub Enterprise Server. Manipulations will take place in `github.com` repositories if you didn't set this variable.

`--dry-run` for all operations are supported.

### Create Labels

```console
$ takolabel create
```

Create labels to the repositories specified in the YAML file.

Write labels settings in `takolabel_create.yml` and put in the same directory as the one where you run the command.

e.g.

```yaml
repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
labels:
  - name: Label 1
    description: This is the label one 
    color: ff0000
  - name: Label 2
    description: This is the label two
    color: 00ff00
  - name: Label 3
    description: This is the label three
    color: 0000ff
```

### Delete Labels

```console
$ takolabel delete
```

Delete labels in the repositories specified in the YAML file.

Write labels settings in `takolabel_delete.yml` and put in the same directory as the one where you run the command.

e.g.

```yaml
repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
labels:
  - Label 1
  - Label 2
  - Label 3
```

### Sync Labels

```console
$ takolabel sync
```

Sync labels in the repositories to those of specified in the YAML file.

Write labels settings in `takolabel_sync.yml` and put in the same directory as the one where you run the command.

e.g.

```yaml
repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
labels:
  - name: Label 1
    description: This is the label one
    color: ff0000
  - name: Label 2
    description: This is the label two
    color: 00ff00
  - name: Label 3
    description: This is the label three
    color: 0000ff
```

### Empty Labels

```console
$ takolabel empty
```

Delete all the labels in the repositories specified in the YAML file.

Write repositories settings in `takolabel_empty.yml` and put in the same directory as the one where you run the command.

e.g.

```yaml
repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
```
