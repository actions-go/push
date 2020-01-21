<p align="center">
  <a href="https://github.com/actions/typescript-action/actions"><img alt="typescript-action status" src="https://github.com/actions/typescript-action/workflows/build-test/badge.svg"></a>
</p>

# Create an Action using Golang

Use this template to bootstrap the creation of a Golang action.:rocket:

This template includes compiliation support, tests, a validation workflow, publishing, and versioning guidance.  

## Create an action from this template

Click the `Use this Template` and provide the new repo details for your action

## Code in Master

Install the dependencies  
```bash
$ go mod download
```

Build the code
```bash
$ go build
```

Run the tests :heavy_check_mark:  
```bash
$ go test -v ./...

  === RUN   TestRunMain
  --- PASS: TestRunMain (0.01s)
  PASS
  ok      github.com/actions-go/go-action    0.315s

...
```

Commit the pre-built binaries to `./dist/main_linux`, `./dist/main_darwin` and `./dist/main_windows.exe` for respectively linux, mac OS and windows actions.

## Change action.yml

The action.yml contains defines the inputs and output for your action.

Update the action.yml with your name, description, inputs and outputs for your action.

See the [documentation](https://help.github.com/en/articles/metadata-syntax-for-github-actions)

<!---
## Change the Code

Most toolkit and CI/CD operations involve async operations so the action is run in an async function.

```javascript
import * as core from '@actions/core';
...

async function run() {
  try { 
      ...
  } 
  catch (error) {
    core.setFailed(error.message);
  }
}

run()
```

See the [toolkit documentation](https://github.com/actions/toolkit/blob/master/README.md#packages) for the various packages.

-->

## Publish to a distribution branch

Actions are run from GitHub repos.  We will create a releases branch and only checkin production modules (core in this case). 

```bash
$ git checkout -b releases/v1
$ git commit -a -m "prod dependencies"
```

```bash
$ git commit -a -m "prod dependencies"
$ git push origin releases/v1
```

Your action is now published! :rocket: 

See the [versioning documentation](https://github.com/actions/toolkit/blob/master/docs/action-versioning.md)

<!--
## Validate

You can now validate the action by referencing the releases/v1 branch

```yaml
uses: actions/typescript-action@releases/v1
with:
  milliseconds: 1000
```

See the [actions tab](https://github.com/actions/javascript-action/actions) for runs of this action! :rocket:

-->

## Usage:

After testing you can [create a v1 tag](https://github.com/actions/toolkit/blob/master/docs/action-versioning.md) to reference the stable and tested action

```yaml
uses: actions-go/go-action@master
with:
  milliseconds: 1000
```
