#!/bin/bash

# standard bash error handling
set -o errexit;
set -o pipefail;
set -o nounset;
# debug commands
set -x;

KUBECTL_VERSION=v1.17.0

install_kfx() {
    make build
}

install_kubectl() {
    # Install needed executables
    wget https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    export PATH=kubectl:${PATH}
}

main() {
    install_kubectl
    install_kfx
    mv bin/kfx-linux kfx
    chmod +x ./kfx
    ./kfx install kf

    # TODO: invoke noteboks, pipelines, operator tests here
    # teardown will happen automatically on exit; GitHub Actions does this for us
}

main
