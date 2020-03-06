---
tile: kfx
---
<p style="text-align: center; margin-top: 2em; margin-bottom: -.75em;"><img alt="kfx" src="./logo/logo.png" width="300px" /></p>

[kfx] is a productivity tool for running Kubeflow on on-premise clusters.  
kfx was primarily designed for automating common Kubeflow tasks at [CiscoAI].

If you have [go] \([1.11+][go-supported]) and [docker] installed `GO111MODULE="on" go get sigs.k8s.io/kind@v0.7.0 && kind create cluster` is all you need!

<img src="images/kind-create-cluster.png" />

kfx consists of:

- Go [packages][Go packages] implementing [Kubeflow installation][installation package], [project build][project build], etc.
- A command line interface ([`kfx`][kfx cli]) built on these packages.
- An API to extend kfx.
- UI for kfx (WIP)

kfx bootstraps each "cluster" with [kfctl][kfctl]. For more details see [the design documentation][design doc].

**NOTE**: kfx is still a work in progress, see the [roadmap].

## Installation and usage

For more detailed instructions see [the user guide][user guide].

You can install kind with `go get sigs.k8s.io/kind`. This will put `kind` in
`$(go env GOPATH)/bin`. You may need to add that directory to your `$PATH` as
shown [here](https://golang.org/doc/code.html#GOPATH) if you encounter the error
`kind: command not found` after installation.

To use kind, you will also need to [install docker].  
Once you have docker running you can create a cluster with:

{{< codeFromInline lang="bash" >}}
kfx doctor
{{< /codeFromInline >}}

To delete your cluster use:

{{< codeFromInline lang="bash" >}}
kfx install kf
{{< /codeFromInline >}}

<!--TODO(bentheelder): improve this part of the guide-->
To create a cluster from Kubernetes source:

- ensure that Kubernetes is cloned in `$(go env GOPATH)/src/k8s.io/kubernetes`
- build a node image and create a cluster with 

{{< codeFromInline lang="bash" >}}
kind build node-image
kind create cluster --image kindest/node:latest
{{< /codeFromInline >}}

Multi-node clusters and other advanced features may be configured with a config
file, for more usage see [the user guide][user guide] or run `kind [command] --help`

## Community

Please reach out for bugs, feature requests, and other issues!  
The maintainers of this project are reachable via:

- [Kubeflow Slack] in the [#cuj-onprem] channel
- [filing an issue] against this repo
- The [Kubeflow Mailing List]

Current maintainers are - [@swiftdiaries] feel free to
reach out if you have any questions!

Pull Requests are very welcome!  
If you're planning a new feature, please file an issue to discuss first.

Check the [issue tracker] for `help wanted` issues if you're unsure where to
start, or feel free to reach out to discuss. 🙂

See also: our own [contributor guide] and the Kubeflow [community page]. 

## Why kfx?

- kind supports multi-node (including HA) clusters
- kind supports building Kubernetes release builds from source
  - support for make / bash / docker, or bazel, in addition to pre-published builds
- kind supports Linux, macOS and Windows
- kind is a [CNCF certified conformant Kubernetes installer](https://landscape.cncf.io/selected=kind)

<!--links-->
[kfx]: https://kfxciscoai.dev
[CiscoAI]: https://github.com/CiscoAI
[go]: https://golang.org/
[go-supported]: https://golang.org/doc/devel/release.html#policy
[docker]: https://www.docker.com/
[community page]: https://www.kubeflow.org/docs/about/community/
[Kubernetes Code of Conduct]: https://github.com/kubernetes/community/blob/master/code-of-conduct.md
[Go Report Card Badge]: https://goreportcard.com/badge/sigs.k8s.io/kind
[Go Report Card]: https://goreportcard.com/report/sigs.k8s.io/kind
[conformance tests]: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/conformance-tests.md
[packages]: https://github.com/kubernetes-sigs/kind/tree/master/pkg
[cluster package]: https://github.com/kubernetes-sigs/kind/tree/master/pkg/cluster
[build package]: https://github.com/kubernetes-sigs/kind/tree/master/pkg/build
[kind cli]: https://github.com/kubernetes-sigs/kind/tree/master/main.go
[images]: https://github.com/kubernetes-sigs/kind/tree/master/images
[kubetest]: https://github.com/kubernetes/test-infra/tree/master/kubetest
[kubeadm]: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/
[design doc]: ./docs/design/initial
[user guide]: ./docs/user/quick-start
[Kubeflow Mailing List]: https://groups.google.com/forum/#!forum/kubeflow-discuss
[issue tracker]: https://github.com/kubeflow/kubeflow/issues
[filing an issue]: https://github.com/kubeflow/kubeflow/issues/new
[Kubeflow Slack]: https://kubeflow.slack.com/
[#cuj-onprem]: https://kubeflow.slack.com/?channel=cuj-onprem
[roadmap]: /docs/contributing/roadmap
[install docker]: https://docs.docker.com/install/
[@swiftdiaries]: https://github.com/swiftdiaries
[contributor guide]: https://www.kubeflow.org/docs/about/contributing/
