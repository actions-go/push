# Push changes on your repository

Use this action to push changes done by other actions to the upstream repository

## Inputs

### `author-email`

The email that will appear in commits when changes needs to be committed.

### `author-name`

The name that will appear in commits when changes needs to be committed.

### `create-commit`

Instructs to create a commit with changed files.

### `commit-message`

The commit message used when changes needs to be committed.

### `ref`

The name of the local reference to push.

### `remote`

The name of the remote on which to push the changes. Defaults to the tracked remote

### `remote-ref`

The name of the remote reference to pushto. Defaults to the tracked remote branch.

### `fail-if-empty`

Fail the action in case there is nothing to do.

## Outputs

### `empty`

TRUE when the action did not perform anything.


## Example usage


```yaml
uses: actions-go/push@master
with:
  commit-message: '[Auto] update pre-puilt dist packages'
  remote: origin
```