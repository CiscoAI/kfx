# Getting Started

This doc gives a quick-start on using the `create-kf-app` tool.

## Quick Start

Ensure you have the latest release from [here](https://github.com/CiscoAI/create-kf-app/releases).
Download the right binary, unpack and place it in your `$PATH`.

```bash
    tar -zvxf create-kf-app-<platform>.tar.gz
    export PATH=$PATH:"<path to kfctl>"
    create-kf-app-<platform> init --name kf
```

This creates a local cluster with KinD (Kubernetes in Docker). With the `KUBECONFIG` placed at `~/.kube/kind-config/kf-kind`.

Set the KUBECONFIG variable.

`export $KUBECONFIG=~/.kube/kind-config/kf-kind`

Now, you can check the installed Kubeflow components using `kubectl get po -n kubeflow`.
You should be all components running fine except the profile-controller deployment. Since this is just a single user cluster, that's fine.

To connect to the Kubeflow UI:

`kubectl port-forward svc/ambassador 8080:80`

And click on this [link](localhost:8080).
