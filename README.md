# cello

create-kf-app is a tool to package ML apps for Kubernetes.

## What is create-kf-app?

- Micro-PaaS for ML on Kubernetes.
- A tool to structure and organize your Kubeflow apps.
- Internally, it uses [KinD](#whykind) to create a k8s cluster on your local machine.

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
