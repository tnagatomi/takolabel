# takolabel

## Installation

### Mac

```console
$ brew install tommy6073/tap/takolabel
```

### Other platforms

Download from [Releases page](https://github.com/tommy6073/takolabel/releases) in this repository.

## Usage

Set variables below in `takolabel.env` and put in the same directory as the one where you run the command.

- GITHUB_TOKEN
  - A token with `repo` scope is needed if it will be run on a private repository. `public_repo` scope will suffice if it's a public repository.
- GITHUB_SERVER_URL (e.g. `https://ghe.example.com/`) (optional)
  - Set this variable if you want to work with repositories hosted on GitHub Enterprise server. Manipulations will take place in `github.com` repositories if you didn't set this variable.

`--dry-run` for all operations are supported.

### Create Labels

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

Run command.

```console
$ takolabel create
```

### Delete Labels

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

Run command (you will be confirmed).

```console
$ takolabel delete
```
