### create-kf-app
----------------

create-kf-app is a tool to accelerate and simplify ML app development on Kubernetes.

## Installation and usage

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

Install from Source
-------------------

**NOTE**: please use the latest go to do this, ideally go 1.12.7 or greater.

```bash
git clone https://github.com/CiscoAI/create-kf-app
export GO111MODULE=on
cd create-kf-app
make build
```

To use create-kf-app, you will need to:

* install docker
* install kubectl

Usage
-----
Once you have docker and kubectl you can create a cluster with Kubeflow by,

```bash
opsys=macos # linux
cd bin/
chmod +x ./create-kf-app-${opsys}
./create-kf-app init --name kf
```

To delete your cluster use `create-kf-app delete cluster --name kf`

Current maintainers are [@swiftdiaries](github.com/swiftdiaries) - feel free to reach out if you have any questions!

## Why kind?

- kind supports multi-node (including HA) clusters
- kind is written in go, can be used as a library, has stable releases
- kind supports Windows in addition to MacOS and Linux
- kind is a [CNCF certified conformant Kubernetes installer](https://landscape.cncf.io/selected=kind)

## Alternatives

Some other open source projects with slightly different but overlapping use cases, features etc.

- https://www.arrikto.com/minikf/
- https://github.com/bsycorp/kind
- https://github.com/ubuntu/microk8s
- https://github.com/kinvolk/kube-spawn
- https://github.com/kubernetes/minikube
- https://github.com/danderson/virtuakube
- https://github.com/kubernetes-sigs/kubeadm-dind-cluster

