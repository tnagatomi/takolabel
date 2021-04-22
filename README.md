# takolabel

## Installation

### Mac

```console
$ brew install tommy6073/tap/takolabel
```

### Other platforms

Download from [Releases page](https://github.com/tommy6073/takolabel/releases) in this repository.

## Usage

`takolabel.env` に以下の変数を設定します

* BASE_URL (例: `https://ghe.example.com/`)
  * 設定しなかった場合は `github.com` に対して操作を行います。
* GITHUB_TOKEN

`takolabel_create.yaml` にラベルの設定を書いて実行します。

Example:

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
