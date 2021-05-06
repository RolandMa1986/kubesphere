# External Repository Staging Area

This directory is the staging area for packages that have been split to their
own repository. The content here will be periodically published to respective
top-level kubesphere.io repositories.

Repositories currently staged here:

- [`kubesphere.io/client-go`](https://github.com/kubesphere/client-go)


The code in the staging/ directory is authoritative, i.e. the only copy of the
code. You can directly modify such code.

## Using staged repositories from KubeSphere code

KubeSphere code uses the repositories in this directory via symlinks in the
`vendor/kubesphere.io` directory into this staging area. For example, when
KubeSphere code imports a package from the `kubesphere.io/client-go` repository, that
import is resolved to `staging/src/kubesphere.io/client-go` relative to the project
root:

```go
// pkg/example/some_code.go
package example

import (
  "kubesphere.io/client-go/" // resolves to staging/src/kubesphere.io/client-go/dynamic
)
```

Once the change-over to external repositories is complete, these repositories
will actually be vendored from `kubesphere.io/<package-name>`.

For other Go applictions, you can install the module with 

```go
go get kubesphere.io/client-go@v0.3.1
```

## Creating a new repository in staging

### Adding the staging repository in `kubesphere/kubesphere`:

1. Add a proposal to sig-architecture in [community](https://github.com/kubesphere/community/). Waiting for approval to creating the staging repository.

2. Once approval has been granted, create the new staging repository.

3. Add a symlink to the staging repo in `vendor/kubesphere.io`.

4. Add all mandatory template files to the staging repo such as README.md, LICENSE, OWNER, CONTRIBUTING.md.


### Creating the published repository

1. Create a repository in the KubeSphere org. The published repository **must** have an
initial empty commit.

2. Setup branch protection and enable access to the `ks-publishing-bot` bot.

3. Once the repository has been created in the KubeSphere org, update the publishing-bot to publish the staging repository by updating:

    - [`rules.yaml`](/staging/publishing/rules.yaml):
    Make sure that the list of dependencies reflects the staging repos in the `Godeps.json` file.

4. Add the repo to the list of staging repos in this `README.md` file.

## Creating new release branchs and tags for the target repo

### Updating branch mapping rules

When a new release branch has been created in the kubesphere repository. [`rules.yaml`](/staging/publishing/rules.yaml) is needed to be modified. Add the new branch mapping rules for all of the existing repos.

```
rules:
- destination: client-go
  library: true
  branches:
  - source:
      branch: master
      dir: staging/src/kubesphere.io/client-go
    name: master
  - source:
      branch: release-3.1
      dir: staging/src/kubesphere.io/client-go
    name: release-3.1
```
### Creating the release tags

The publishing-bot don't support creating tags yet. So we have to create tag for the target repo mannully. The go modules is using the v0.x.y tags because all the client-go tags greater than 1 (v2, v3, etc) are all incompatible modules. Such as：

```
kubesphere vx.y.z -> client-go v0.x.y
kubesphere v3.1.0 -> client-go v0.3.1
kubesphere v4.0.0 -> client-go v0.4.0
```