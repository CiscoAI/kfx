# Getting Started

This doc gives a quick-start on using the `kfx` tool.

## Quick Start

Ensure you have the latest release from [here](https://github.com/CiscoAI/kfx/releases).
Download the right binary, unpack and place it in your `$PATH`.

```bash
    export PATH=$PATH:"<path to kfctl>"
    kfx create cluster
    kfx install kf
```

Set the KUBECONFIG variable.

`export $KUBECONFIG=~/.kube/kind-config/kf-kind`

Now, you can check the installed Kubeflow components using `kubectl get po -n kubeflow`.

To connect to the Kubeflow UI:

`kfx ui kf`

And click on this [link](localhost:8080).
