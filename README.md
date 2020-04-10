# kfx

kfx is a productivity tool for Kubeflow on-premise.

🚀 on-premise ML development with Kubeflow.

Table of Contents
=================

   * [What is kfx?](#what-is-kfx)
      * [kfx CLI](#kfx-cli)
      * [kfx API](#kfx-api)
      * [kfx UI](#kfx-ui)
   * [Installation](#installation)
   * [Questions/Feedback](#faq)

---

## <a name="what-is-kfx"></a>What is kfx?

- A tool to perform common Kubeflow tasks.
- Opinionated way of using Kubeflow.

### <a name="kfx-cli"></a> kfx CLI - Presto

A very fast way to use Kubeflow.

#### Who is this for?

- People who prefer CLIs over UIs.
- ML Engineers, Data Scientists.

### <a name="kfx-api"></a> kfx API - Andante

A API-centric way to use Kubeflow.

- People who want to extend kfx for custom tasks.
- ML Engineers.


### <a name="kfx-ui"></a> kfx UI - Allegro

A UI-centric way to use Kubeflow.

#### Who is this for?

- People who prefer UIs over CLIs.
- ML Engineers, Data Scientists.

## <a name="installation"></a> Installation and usage

### Requirements

To use kfx, you will need to:

- install docker
- install kubectl

If you have Docker and kubectl installed on your machine, you are 💯 to go

### Installation

Download the binaries:

```bash
export KFX_VERSION=v0.1.0-alpha
export OPSYS=linux # or darwin
curl http://storage.googleapis.com/kfx-releases/${KFX_VERSION}/${OPSYS}/kfx > kfx && chmod +x kfx
```
Add the binary to your $PATH.

```bash
kfx install kf --name starter-pack
```

And you're ready to Kubeflow !

## <a name="faq"></a> Questions / Feedback

Current maintainer [@swiftdiaries](github.com/swiftdiaries) - feel free to reach out if you have any questions or discuss Kubeflow on-prem.
