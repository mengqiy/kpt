// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package tutorials

var BuildingSolutionsShort = `How to build solutions using kpt with other tools from the ecosystem.`
var BuildingSolutionsLong = `
Also see [Building Solutions with kpt](../README.md#building-solutions-with-kpt)

### Packaging: ` + "`" + `kpt get` + "`" + `, ` + "`" + `kpt update` + "`" + `

  Packaging enables fully or partially specified Resource configuration
  + related artifacts to be published and consumed, as well as facilitates
  updating configuration from upstream.

  Example Use Cases:

  - Fetch a *Blueprint* or *Example* and fork or extend it
  - Fetch *Configuration Function* Resources
  - Fetch configuration to be applied directly to a cluster

  - Update a forked *Blueprint* from upstream
  - Update a *Configuration Function* Resource from upstream
  - Update configuration applied to a cluster

  Example:

  Fetch a Blueprint:

    kpt get https://github.com/kubernetes/examples/cassandra cassandra/

  Update a Blueprint to a specific git commit, merging Resource updates with
  local changes to the fork:
  
    kpt update cassandra@322d78b --strategy resource-merge 

### Development: ` + "`" + `kustomize build` + "`" + `, ` + "`" + `kustomize config run` + "`" + `

  Development of configuration is about developing the configuration which will
  be applied to an apiserver.

  It may involve a number of activities:

  1. Developing Abstractions

     Abstractions allow a higher-level or more specialized API to be defined
     which may generate other Resources. 

     - Templating Resources -- Jinja, YTT, Helm
     - Generating Resources From DSLs --Cue,  Ksonnet, Jsonnet, Terraform
     - Generating Resources Programmatically -- Starlark, TypeScript
  
  2. Developing Blueprint Customizations

     Blueprints allow low-level Resource configuration to be published and
     customized for a specialized case.

     - Change replica counts
     - Change container image
     - Add environment variables

  3. Developing Variant Customizations

     Variants apply customizations for a specific environment.  They
     are useful when the same package needs to be run in several environments,
     but with different opinions. 

     - Dev, Test, Staging, Production
     - us-west, us-east, us-central, asia, europe

  4. Injecting Cross-Cutting elements into Resources

     Injection is useful for applying policies or cross-cutting logic to
     a collection of Resources.  Notably, the injector may be loosely
     coupled from the package being injected.

     - T-Shirt sizing containers based on annotations
     - Injecting side car containers
     - Injecting init containers

  5. Validating Resources

     Validating Resources is important for applying linting or organizational
     opinions.

     - Ensuring resource reservations are specified
     - Ensuring container images are tagged

  How does kpt facilitate these?

  The kpt architecture enables decoupling programs and tools from
  the packaged configuration itself by applying functions (encapsulated in containers)
  to the local configuration.
  That is the packages themselves contain Resource configuration
  rather than code (e.g. templates, DSLs, etc).  The packaged Resources may
  be modified or expanded by external programs, such as ` + "`" + `kustomize` + "`" + `.

  ` + "`" + `kustomize` + "`" + ` is a tool which can be used to develop configuration by:

   - defining customization variants
   - applying functions which may be used for developing abstractions, cross-cutting
     modifications, and validation

  Example Use Cases:

  - Develop variants for test, staging, production versions of config
  - Develop a high-level "App" abstraction API which takes only a few inputs
    and generates a Deployment, Service and ConfigMap
  - Develop an annotation for t-shirt sizing resource reservations, setting cpu
    and memory to values for small, medium, and large
  - Develop validation to ensure container images are always tagged

  Examples:

  See the [example functions](https://github.com/kubernetes-sigs/kustomize/tree/master/functions/examples)

    kustomize config run DIR

  See the [kustomize examples](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)

    kustomize build DIR  

### Actuation: ` + "`" + `kubectl apply` + "`" + `, ` + "`" + `kustomize status` + "`" + `, ` + "`" + `kustomize prune` + "`" + `

  Applying a collection of configuration may involve several steps, and may require
  orchestrating the actuation of several different packages.  The building blocks of
  actuating configuration are:

  1. Apply

     - Take collection of local Resources and send to the cluster
     - Merge locally defined desired state with cluster defined desired state
       (e.g. keep replica count defined by autoscaler in the cluster)

  2. Status

     - Track the status of the changes until they have been fully rolled out
     - Block until the process completes, or fails to make progress for some period of time 
       (e.g. timesout)

  3. Prune

     - Identify Resources that exist in the cluster, but have been deleted locally and delete them
     - Support diff / dry-run

  The kpt architecture facilitates using the Kubernetes project based tooling,
  such as ` + "`" + `kubectl` + "`" + ` and ` + "`" + `kustomize` + "`" + ` for actuating configuration changes.

  Example Use Cases:

  1. Apply a package of configuration to a cluster
  2. Wait until it is successful, printing an error on failure
  3. Delete Resources that have deleted from the package since it was last applied

  **Note:** the actuation steps may be performed by automation using a GitOps approach --
  e.g. trigger Google Cloud Build to perform the actuation when PRs are merged into
  release branches.

  Examples:

    # apply non-local Resources -- skips config-functions
    kustomize config cat DIR | kubectl apply -f -
 
    # block on completion of changes
    kustomize status

    # delete Resources removed from the package
    kustomize prune

### Visibility / Inspection: ` + "`" + `kustomize config tree` + "`" + `, ` + "`" + `kustomize config grep` + "`" + `

  When working with configuration as data, the configuration may become verbose.
  This makes it challenging to quickly understand the state of the system declared
  locally.
  
  Tools such as ` + "`" + `kustomize config tree` + "`" + ` help parse and visualize packages of configuration.
  They may be used with tools such as ` + "`" + `kustomize config grep` + "`" + ` to query configuration.
  
  Example Use Cases:
  
  - Display all the Resources in a package
  - Display all Resources in a package containing an untagged container image
  - Display all Resources in a package containing a container without resource reservations

  Examples:

    # display resources, as well as container names and images
    kustomize config tree DIR/ --name --image

    # find Resources named nginx
    kustomize config grep "metadata.name=nginx" my-dir/    

### Putting It All Together

1. Fetch a package of configuration

       kpt get https://github.com/kubernetes/examples/cassandra cassandra/

2. View the package Resources

       kustomize config tree cassandra/ --image --name

3. Customize or Develop the package

      # view "setters" for the package
      kustomize config set cassandra/

       # add configuration functions, then run
       kustomize config run cassandra/

       # or add kustomize variants that use it as a base
       mkdir prod/
       vi prod/Kustomization.yaml

4. Apply the package to a cluster

       kustomize apply cassandra/`

var FaqShort = `Frequently Asked Questions`
var FaqLong = `
Q: **How does kpt fit in with the Kubernetes ecosystem?**
A: ` + "`" + `kpt` + "`" + ` is intended to be composed with other tools from the Kubernetes ecosystem.
   Rather than attempting to solve all problems related to configuration, kpt
   is focused solving how to publish and consume configuration packages.  ` + "`" + `kpt` + "`" + `
   was developed to complement both the OSS Kubernetes project tools such as
   ` + "`" + `kubectl` + "`" + ` and ` + "`" + `kustomize` + "`" + `, and other tools developed as part of the broader
   ecosystem.

Q: **How do I use kpt with the Kubernetes project tools?**
A: ` + "`" + `kpt` + "`" + ` may be used to publish, fetch and update configuration packaging.
   The project tools may be used to manipulate and apply the fetched configuration.

Q: **How can I use kpt to create blueprint packages?**
A: Blueprints may be published as ` + "`" + `kpt` + "`" + ` packages, using kpt for fetching and
   updating the packages from upstream.
   The local copy of the blueprint package may be directly edited, or other
   customization techniques may be applied (e.g. ` + "`" + `kustomize build` + "`" + `).

Q: **What are some examples of blueprint customization techniques?**
A: - Using configuration functions (i.e. ` + "`" + `kustomize config run` + "`" + `)
   - Using Kustomizations (i.e. ` + "`" + `kustomize build` + "`" + `)
   - Duck-typed setter commands (i.e. ` + "`" + `kustomize duck CMD` + "`" + `)

Q: **What are configuration functions?**
A: Configuration functions are programs applied to configuration which may generate, transform
   or validate new or existing configuration.  Functions are typically published as container
   images applied to a configuration package.
   See ` + "`" + `kustomize help config run` + "`" + ` for more information.

Q: **What are setters?**
A: Setters are similar to the imperative ` + "`" + `kubectl set` + "`" + ` commands, but operate against local
   Resource configuration and are configured as annotations on Resource fields.
   See ` + "`" + `kustomize help config set` + "`" + ` for more information.

Q: **How do I parameterize kpt packages?**
A: For performing simple substitutions, ` + "`" + `kustomize config set` + "`" + ` may be used to replace
   marker values with values provided on the commandline.
   Alternatively, substitutions may be externalized from the package using configuration functions
   which can generate or transform configuration.
   See ` + "`" + `kustomize help config set` + "`" + ` and ` + "`" + `kustomize help config run` + "`" + ` for more information.

Q: **Does kpt work for non-Resource packages such as Terraform or Helm Charts?**
A: Sorta, kpt packages can contain non-Resource packaging artifacts.  These
   artifacts do not support Resource specific operations -- e.g.
   The ` + "`" + `update` + "`" + ` command ` + "`" + `resource-merge` + "`" + ` strategy will not work against them,
   but the ` + "`" + `alpha-git-patch` + "`" + ` strategy will.
   
Q: **How can I apply kpt packages to a cluster?**
A: kpt is designed to work with the OSS Kubernetes Kubernetes tools such
   as ` + "`" + `kubectl` + "`" + ` and ` + "`" + `kustomize` + "`" + ` -- e.g. ` + "`" + `kustomize apply PKG/` + "`" + ` or
   ` + "`" + `kustomize config cat PKG/ | kubectl apply -f -` + "`" + `.
   Use ` + "`" + `kustomize config cat` + "`" + ` so that only non-config function Resources are applied.

`

var FetchAPackageShort = `How to fetch a package from a remote source`
var FetchAPackageLong = `
### Synopsis

Packages are directories of Configuration published as subdirectories to git repositories.

- No additional package metadata or structure is required for a package to be fetched
- Format is natively compatible with ` + "`" + `kubectl apply` + "`" + ` and ` + "`" + `kustomize` + "`" + `
- May be fetched and updated to specific revisions (using git tags or branches)
- May also include non-configuration files or metadata

### Fetch the Cassandra package

  Fetch a "raw" package (e.g. config only -- no kpt metadata) from the kubernetes examples repo.

	kpt get  https://github.com/kubernetes/examples/cassandra cassandra/

  ` + "`" + `kpt get` + "`" + ` fetched the remote package from HEAD of the
  https://github.com/kubernetes/examples master branch.

	$ kustomize config tree cassandra/
	cassandra
	├── [cassandra-service.yaml]  v1.Service cassandra
	├── [cassandra-statefulset.yaml]  apps/v1.StatefulSet cassandra
	└── [cassandra-statefulset.yaml]  storage.k8s.io/v1.StorageClass fast
	
  ` + "`" + `kustomize config tree` + "`" + ` printed the package structure -- displaying both the Resources as well as the
  files the Resources are specified in.

	$ kpt desc cassandra
	+-----------------+-----------+----------------------------------------+-----------+---------+---------+
	| LOCAL DIRECTORY |   NAME    |           SOURCE REPOSITORY            |  SUBPATH  | VERSION | COMMIT  |
	+-----------------+-----------+----------------------------------------+-----------+---------+---------+
	| cassandra       | cassandra | https://github.com/kubernetes/examples | cassandra | master  | 1543966 |
	+-----------------+-----------+----------------------------------------+-----------+---------+---------+

  ` + "`" + `kpt desc LOCAL_PACKAGE` + "`" + ` prints information about the source of the package -- e.g. 
  the repo, subdirectory, etc.

### Fetch the Guestbook package

	$ kpt get https://github.com/kubernetes/examples/guestbook ./my-guestbook-copy

  The guestbook package contains multiple guest book instances in separate
  subdirectories.

	$ kustomize config tree my-guestbook-copy/
	my-guestbook-copy
	├── [frontend-deployment.yaml]  apps/v1.Deployment frontend
	├── [frontend-service.yaml]  v1.Service frontend
    ...
	├── all-in-one
	│   ├── [frontend.yaml]  apps/v1.Deployment frontend
	│   ├── [frontend.yaml]  v1.Service frontend
	│   ├── [guestbook-all-in-one.yaml]  apps/v1.Deployment frontend
    ...
	└── legacy
		├── [frontend-controller.yaml]  v1.ReplicationController frontend
		├── [redis-master-controller.yaml]  v1.ReplicationController redis-master
		└── [redis-slave-controller.yaml]  v1.ReplicationController redis-slave

  The separate guestbook subpackages contain variants of the same guestbook application.
  To fetch only the all-in-one instance, specify the instance subdirectory as
  part of the package.

	$ kpt get https://github.com/kubernetes/examples/guestbook/all-in-one ./new-guestbook-copy

  ` + "`" + `kpt get` + "`" + ` only fetched the all-in-one subpackage.

	$ kustomize config tree new-guestbook-copy
	new-guestbook-copy
	├── [frontend.yaml]  apps/v1.Deployment frontend
	├── [frontend.yaml]  v1.Service frontend
	├── [guestbook-all-in-one.yaml]  apps/v1.Deployment frontend
	├── [guestbook-all-in-one.yaml]  v1.Service frontend
	├── [guestbook-all-in-one.yaml]  apps/v1.Deployment redis-master
	├── [guestbook-all-in-one.yaml]  v1.Service redis-master
	├── [guestbook-all-in-one.yaml]  apps/v1.Deployment redis-slave
	├── [guestbook-all-in-one.yaml]  v1.Service redis-slave
	├── [redis-slave.yaml]  apps/v1.Deployment redis-slave
	└── [redis-slave.yaml]  v1.Service redis-slave

### Package Versioning

  Since packages are stored in git, git references may be used to fetch a specific version
  of a package.

	kpt get https://github.com/GoogleContainerTools/kpt/package-examples/hello-world@v0.1.0 hello-world/

  Specifying '@version' after the package uri fetched the package at that revision.
  The version may be a git branch, tag or ref.
  
  Note: git references may also be used with ` + "`" + `kpt update` + "`" + ` to rollout new configuration versions.
  See ` + "`" + `kpt help update` + "`" + ` for more information.

### New Package From Kustomize Output

  ` + "`" + `kpt get` + "`" + ` may also be used to convert ` + "`" + `kustomize` + "`" + ` output into a package

    # fetch a kustomize example
	kpt get https://github.com/kubernetes-sigs/kustomize/examples/wordpress wordpress/
	
	# build the kustomize package and use ` + "`" + `kpt get` + "`" + ` to write the output to a directory
	kustomize build wordpress/ | kpt get - wordpress-expanded/

  This expanded the Kustomization into a new package

	$ kustomize config tree wordpress-expanded/
	wordpress-expanded
	├── [demo-mysql-pass_secret.yaml]  v1.Secret demo-mysql-pass
	├── [demo-mysql_deployment.yaml]  apps/v1beta2.Deployment demo-mysql
	├── [demo-mysql_service.yaml]  v1.Service demo-mysql
	├── [demo-wordpress_deployment.yaml]  apps/v1beta2.Deployment demo-wordpress
	└── [demo-wordpress_service.yaml]  v1.Service demo-wordpress

### New Package From Helm Output

  ` + "`" + `kpt get` + "`" + ` may be used to write expanded ` + "`" + `helm` + "`" + ` templates to packages.

	helm fetch stable/redis
	helm template redis-9.* | kpt get - ./redis-9/

  This imported the expanded package Resources from stdin and created a local kpt package.

	$ kustomize config tree redis-9/
	redis-9
	├── [release-name-redis-headless_service.yaml]  v1.Service release-name-redis-headless
	├── [release-name-redis-health_configmap.yaml]  v1.ConfigMap release-name-redis-health
	├── [release-name-redis-master_service.yaml]  v1.Service release-name-redis-master
	├── [release-name-redis-master_statefulset.yaml]  apps/v1beta2.StatefulSet release-name-redis-master
	├── [release-name-redis-slave_service.yaml]  v1.Service release-name-redis-slave
	├── [release-name-redis-slave_statefulset.yaml]  apps/v1beta2.StatefulSet release-name-redis-slave
	├── [release-name-redis_configmap.yaml]  v1.ConfigMap release-name-redis
	└── [release-name-redis_secret.yaml]  v1.Secret release-name-redis

  The names of the Resource files may be configured using the --pattern flag.

	helm fetch stable/redis
	helm template redis-9.* | kpt get - ./redis-9/ --pattern '%n.resource.yaml'
	
  This configured the generated resource file names to be RESOURCENAME.resource.yaml
  instead of RESOURCENAME_RESOURCETYPE.yaml
  Multiple Resources with the same name are put into the same file:

	$ kustomize config tree redis-9/
	redis-9
	├── [release-name-redis-headless.resource.yaml]  v1.Service release-name-redis-headless
	├── [release-name-redis-health.resource.yaml]  v1.ConfigMap release-name-redis-health
	├── [release-name-redis-master.resource.yaml]  v1.Service release-name-redis-master
	├── [release-name-redis-master.resource.yaml]  apps/v1beta2.StatefulSet release-name-redis-master
	├── [release-name-redis-slave.resource.yaml]  v1.Service release-name-redis-slave
	├── [release-name-redis-slave.resource.yaml]  apps/v1beta2.StatefulSet release-name-redis-slave
	├── [release-name-redis.resource.yaml]  v1.ConfigMap release-name-redis
	└── [release-name-redis.resource.yaml]  v1.Secret release-name-redis
	
 Run ` + "`" + `kpt help get` + "`" + ` for more information on --pattern options`

var FutureDevelopmentShort = `List of features planned for the future`
var FutureDevelopmentLong = `
### Synopsis

Following is a list of ` + "`" + `kpt` + "`" + ` features planned for future development:

- Fetch the most recent release by default if it exists
  - Follow go module semantics of fetching semantic versions
    if no version is specified

- Check for updates
  - List packages that may be updated
  - List available updates for a package

- Resource merge conflict resolution strategies
  - Fail on conflict
  - Choose local on conflict`

var PublishAPackageShort = `How to publish a package to a remote source`
var PublishAPackageLong = `
### Synopsis

Any bundle of Resource configuration may be published as a package using ` + "`" + `git` + "`" + `.

` + "`" + `kpt init` + "`" + ` initializes a directory with optional package metadata such as a
package documentation file.

### Examples

    kpt init my-package/ --name my-package --description 'fun new package'
    git add my-package && git commit -m 'new kpt package'
    git push origin master`

var UpdateALocalPackageShort = `How to update a package version from a remote source`
var UpdateALocalPackageLong = `
### Synopsis

  Local packages may be updated by rebasing them on top of upstream changes.
  Multiple strategies are supported for merging in updates.
  Run ` + "`" + `kpt help update` + "`" + ` for the list of supported update strategies

  - If no new revision is specified in the update, and the source was a branch, then the package
    will be updated to the tip of that branch.
  - The local package must be committed to git be updated 
  - Updates to packages generated from stdin are not yet supported

## Update an unchanged package

  Prepare the package to be updated

	kpt get https://github.com/GoogleContainerTools/kpt/package-examples/hello-world@v0.1.0 hello-world/
	git add hello-world/ && git commit -m 'fetch hello-world'

  Diff a local package vs a new upstream version

  NOTE: the diff viewer can be controlled by setting KPT_EXTERNAL_DIFF --
  'export KPT_EXTERNAL_DIFF=my-differ'.
  See ` + "`" + `kpt help diff` + "`" + ` for more options.

	kpt diff cockroachdb/@v0.2.0 --diff-type remote
	diff ...
	118c118
	<         image: gcr.io/kpt-dev/hello-world:v0.1.0
	---
	>         image: gcr.io/kpt-dev/hello-world:v0.2.0


  Update the package to the new version.  This requires that the package is unmodified from when
  it was fetched.

	kpt update hello-world@v0.2.0
	git diff hello-world/

  The updates have not been staged by kpt.

## Updating merging remote changes with local changes

  Stage the package to be updated

	kpt get https://github.com/pwittrock/examples/staging/cockroachdb@v1.0.0 cockroachdb/
	git add cockroachdb/ && git commit -m 'fetch cockroachdb'

  Make local edits to the package

	sed -i '' 's/port: 8080/port: 8081/g' ./cockroachdb/cockroachdb-statefulset.yaml
	git add . && git commit -m 'change cockroachdb port from 8080 to 8081'

  Diff the local package vs the original source upstream package -- see what you've changed

	$ kpt diff cockroachdb/
	diff ...
	17c17
	<   - port: 8081
	---
	>   - port: 8080
	50c50
	<   - port: 8081
	---
	>   - port: 8080

  Diff the local package vs a new upstream version -- see what you will be updating to

	$ kpt diff cockroachdb/@v1.4 --diff-type combined
	diff ...
	>     foo: bar
	17c18
	<   - port: 8081
	---
	>   - port: 8080
	50c51
	<   - port: 8081
	---
	>   - port: 8080
	67c68
	<   minAvailable: 67%
	---
	>   minAvailable: 70%
	77c78
	<   replicas: 3
	---
	>   replicas: 7

  Update the package to a new version.

  **NOTE:** --strategy is required when the local package has been changed from its source.
  In this case we have changed the local port field, so we must specify a strategy.

	kpt update cockroachdb@v1.4 --strategy alpha-git-patch
	git diff HEAD^ HEAD

  This merged the upstream changes into the local package, and created a new git commit.

## Update with local merge conflicts

  Stage the package to be updated

	kpt get https://github.com/pwittrock/examples/staging/cockroachdb@v1.0.0 cockroachdb/
	git add cockroachdb/ && git commit -m 'fetch cockroachdb'

  Make local edits to the package.  Edit a field that will be changed upstream.

	kpt cockroachdb set replicas cockroachdb --value 11
	git add . && git commit -m 'change cockroachdb replicas from 3 to 11'

  View the 3way diff -- requires a diff viewer capable of 3way diffs (e.g. meld)

	kpt diff cockroachdb/@v1.4 --diff-type 3way

  This will show that the replicas field cannot be merged without a conflict -- it has
  been changed both in the upstream new package version, and in the local package.

  Go ahead and update the package to a new version anyway.  Expect a merge conflict.

	kpt update cockroachdb@v1.4 --strategy alpha-git-patch

  View the conflict

	$ git diff
	++<<<<<<< HEAD
	 +  replicas: 11
	++=======
	+   replicas: 7
	++>>>>>>> update cockroachdb (https://github.com/pwittrock/examples) from v1.0 (1f356407c2bcd5c56907d366161cbca833679ed1) to v1.4 (a3ea1604962746cd157769ef305951bdd88c628a)

  Fix the conflict and continue with the merge
	
	nano -w ./cockroachdb/cockroachdb-statefulset.yaml
	git add  cockroachdb/
  	git am --continue

  View the updates:

	git diff HEAD^ HEAD

## Manually update by generating a patch

  Stage the package to be updated

	kpt get https://github.com/pwittrock/examples/staging/cockroachdb@v1.0.0 cockroachdb/
	git add cockroachdb/ && git commit -m 'fetch cockroachdb'

  Update the package to a new version.  Expect a merge conflict.

	kpt update cockroachdb@v1.4 --strategy alpha-git-patch --dry-run > patch
	git am -3 --directory cockroachdb < patch

## Update to HEAD of the branch the package was fetched from

  Fetch the package

	kpt get https://github.com/your/repo/here@master here/
	git add cockroachdb/ && git commit -m 'fetch cockroachdb'

  Make upstream changes to the package at https://github.com/your/repo/here on
  the master branch.  Then update it.

	kpt update here/

  This fetched the updates from the upstream master branch.


	tutorials 3-update-a-local-package [flags]`
