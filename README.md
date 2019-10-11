# kfx

kfx is a productivity tool for Kubeflow on-premise.

🚀 Local ML development with Kubeflow.

## What is kfx?

- A tool to perform common Kubeflow tasks from the CLI.
- In-built [KinD](#whykind) support.

### kfx CLI - Presto

A very, very fast way to use Kubeflow.

## Installation and usage

### Requirements

To use kfx, you will need to:

- install docker
- install kubectl

If you have Docker and kubectl installed on your machine, you are 💯 to go

### Installation

Download the binaries:

- Linux 64-bit
`curl https://storage.cloud.google.com/kfx-releases/v0.1.0-alpha/kfx-linux > kfx && chmod +x kfx`

- MacOS 64-bit
`curl https://storage.cloud.google.com/kfx-releases/v0.1.0-alpha/kfx-darwin > kfx && chmod +x kfx`

Download the latest release :arrow_up: :arrow_up:
Add the binary to your $PATH.

```bash
    export PATH=$PATH:"<path to kfctl>"
```

And you're ready to Kubeflow !

To delete your cluster use `kfx delete cluster --name kf`

## TODO: Prestissimo - A Jupyter Kernel for Kubeflow

Note: *even faster than presto*

This creates a Jupyter kernel with the default notebook that kfx installs to connect to a remote Kubernetes cluster.

Reference: [Ciao from Caicloud](https://github.com/caicloud/ciao/blob/master/docs/design.md).

## Questions / Feedback

Current maintainer [@swiftdiaries](github.com/swiftdiaries) - feel free to reach out if you have any questions or discuss Kubeflow on-prem.

## <a name="whykind"></a>Why kind?

- Super light-weight clusters for local development and CI.
- kind supports multi-node (including HA) clusters
- kind supports Windows in addition to MacOS and Linux
- kind is a [CNCF certified conformant Kubernetes installer](https://landscape.cncf.io/selected=kind)

## Alternatives

Some other open source projects with slightly different but overlapping use cases, features etc.

Among this, `miniKF` gives you a Kubeflow ready cluster on top of Minikube.

- [MiniKF](https://www.arrikto.com/minikf/)
- [Bsycorp kind](https://github.com/bsycorp/kind)
- [microk8s](https://github.com/ubuntu/microk8s)
- [kube-spawn](https://github.com/kinvolk/kube-spawn)
- [minikube](https://github.com/kubernetes/minikube)
- [virtualkube](https://github.com/danderson/virtuakube)
- [kubeadm-dind](https://github.com/kubernetes-sigs/kubeadm-dind-cluster)
