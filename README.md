# takolabel

GitHub IssueのLabelを複数指定したリポジトリに一括で作成します。

`takolabel.env` に以下の変数を設定します

* BASE_URL (例: `https://ghe.example.com/`)
* GITHUB_TOKEN

`labels.yaml` にラベルの設定を書いて実行します。

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
