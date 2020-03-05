# kfx

kfx is a productivity tool for Kubeflow on-premise.

🚀 on-premise ML development with Kubeflow.

## What is kfx?

- A tool to perform common Kubeflow tasks.
- Opinionated way of using Kubeflow.

### kfx CLI - Presto

A very fast way to use Kubeflow.

#### Who is this for?

- People who prefer CLIs over UIs.
- ML Engineers, Data Scientists.

### kfx UI - Allegro

A UI-centric way to use Kubeflow.

#### Who is this for?

- People who prefer UIs over CLIs.
- ML Engineers, Data Scientists.

### kfx API - Andante

A API-centric way to use Kubeflow.

- People who want to extend kfx for custom tasks.
- ML Engineers.

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
   kfx install kf --name CiscoAI --pipeline github.com/CiscoAI/bolts-classifier-pipeline
```

And you're ready to Kubeflow !

## Questions / Feedback

Current maintainer [@swiftdiaries](github.com/swiftdiaries) - feel free to reach out if you have any questions or discuss Kubeflow on-prem.
