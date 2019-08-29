### create-kf-app
----------------

create-kf-app is a tool to accelerate and simplify ML app development on Kubernetes.

## Installation and usage

**NOTE**: please use the latest go to do this, ideally go 1.12.7 or greater.

`export GO111MODULE=on`

Do a manual installation by cloning the repo and run `make build` from the repository.

To use create-kf-app, you will need to:

* install docker
* install kubectl

Once you have docker running you can create a cluster with `create-kf-app init --name kf-kind`  

**TODO**: To delete your cluster use `create-kf-app delete cluster --name kf-kind`

Current maintainers are [@swiftdiaries](github.com/swiftdiaries) - feel free to reach out if you have any questions!

## Why kind?

- kind supports multi-node (including HA) clusters
- kind is written in go, can be used as a library, has stable releases
- kind supports Windows in addition to MacOS and Linux
- kind is a [CNCF certified conformant Kubernetes installer](https://landscape.cncf.io/selected=kind)

## Alternatives

Some other open source projects with slightly different but overlapping use cases, features etc.

- https://github.com/bsycorp/kind
- https://github.com/ubuntu/microk8s
- https://github.com/kinvolk/kube-spawn
- https://github.com/kubernetes/minikube
- https://github.com/danderson/virtuakube
- https://github.com/kubernetes-sigs/kubeadm-dind-cluster
- https://www.arrikto.com/minikf/
