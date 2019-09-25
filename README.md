# create-kf-app

create-kf-app is a tool to package ML apps for Kubernetes.

## What is create-kf-app?

- Micro-PaaS for ML on Kubernetes.
- A tool to structure and organize your Kubeflow apps.
- Internally, it uses [KinD](#whykind) to create a k8s cluster on your local machine.

### Presto

A very, very fast way to use Kubeflow.

## Installation and usage

To use create-kf-app, you will need to:

- install docker
- install kubectl

Download the binaries from the release.

🚀 Local ML development with Kubeflow.

If you have Docker and kubectl installed on your machine, you are 💯 to go

Download the latest release.
Unpack the tarball.
Add the binary to your $PATH.

```bash
    tar -zvxf create-kf-app-<platform>.tar.gz
    export PATH=$PATH:"<path to kfctl>"
    create-kf-app-<platform> init --name kf
```

And you're ready to Kubeflow !

To delete your cluster use `create-kf-app delete cluster --name kf`

## kfx

A more advanced user workflow for Kubeflow.

Shares the same goals as the create-kf-app

- Micro-PaaS for ML on Kubernetes.
- A tool to structure and organize your Kubeflow apps.

### Vivace

A lively and fast way to use Kubeflow.

## TODO: Prestissimo - A Jupyter Kernel for Kubeflow

Note: *even faster than presto*

This creates a Jupyter kernel with the default notebook that kfx installs to connect to a remote Kubernetes cluster.

Reference: [Ciao from Caicloud](https://github.com/caicloud/ciao/blob/master/docs/design.md).

## Questions / Feedback

Current maintainer [@swiftdiaries](github.com/swiftdiaries) - feel free to reach out if you have any questions or talk Kubeflow on-prem!

## <a name="whykind"></a>Why kind?

- Super light-weight clusters for local development and CI.
- kind supports multi-node (including HA) clusters
- kind supports Windows in addition to MacOS and Linux
- kind is a [CNCF certified conformant Kubernetes installer](https://landscape.cncf.io/selected=kind)

## Alternatives

Some other open source projects with slightly different but overlapping use cases, features etc.

- [MiniKF](https://www.arrikto.com/minikf/)
- [Bsycorp kind](https://github.com/bsycorp/kind)
- [microk8s](https://github.com/ubuntu/microk8s)
- [kube-spawn](https://github.com/kinvolk/kube-spawn)
- [minikube](https://github.com/kubernetes/minikube)
- [virtualkube](https://github.com/danderson/virtuakube)
- [kubeadm-dind](https://github.com/kubernetes-sigs/kubeadm-dind-cluster)

Among this, `miniKF` gives you a Kubeflow ready cluster on top of Minikube.
