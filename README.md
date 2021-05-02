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
  - A token with `repo` scope will suffice.
- GITHUB_SERVER_URL (e.g. `https://ghe.example.com/`) (optional)
  - Set this variable if you want to work with repositories hosted on GitHub Enterprise server. Manipulations will take place in `github.com` repositories if you didn't set this variable.

Write labels settings in `takolabel_create.yaml` and put in the same directory as the one where you run the command.

e.g.

```yaml
repositories:
  - org: some-org
    repo: some-org-repo-1
  - org: some-org
    repo: some-org-repo-2
  - org: another-org
    repo: another-org-repo-1
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
$ takolabel
```
