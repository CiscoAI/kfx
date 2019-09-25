#!/bin/bash

# standard bash error handling
set -o errexit;
set -o pipefail;
set -o nounset;
# debug commands
set -x;

# working dir to install binaries etc, cleaned up on exit
BIN_DIR="$(mktemp -d)"
# kind binary will be here
KIND="${BIN_DIR}/kind"

# cleanup on exit (useful for running locally)
cleanup() {
    "${KIND}" delete cluster || true
    rm -rf "${BIN_DIR}"
}
trap cleanup EXIT

# util to install the latest kind version into ${BIN_DIR}
install_latest_kind() {
    # clone kind into a tempdir within BIN_DIR
    local tmp_dir
    tmp_dir="$(TMPDIR="${BIN_DIR}" mktemp -d "${BIN_DIR}/kind-source.XXXXX")"
    cd "${tmp_dir}" || exit
    git clone https://github.com/kubernetes-sigs/kind && cd ./kind
    make install INSTALL_DIR="${BIN_DIR}"
    export PATH=${BIN_DIR}:${PATH}
    cd -
}

# util to install a released kind version into ${BIN_DIR}
install_kind_release() {
    VERSION="v0.4.0"
    KIND_BINARY_URL="https://github.com/kubernetes-sigs/kind/releases/download/${VERSION}/kind-linux-amd64"
    wget -O "${KIND}" "${KIND_BINARY_URL}"
    chmod +x "${KIND}"
    export PATH=${KIND}:${PATH}
}

install_kfx() {
    make build
}

install_kubectl() {
    # Install needed executables
    wget https://storage.googleapis.com/kubernetes-release/release/v1.14.0/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    export PATH=kubectl:${PATH}
}

main() {
    # get kind
    install_kind_release
    install_kfx
    mv bin/kfx-linux kfx
    export PATH=kfx:${PATH}
    chmod +x kfx
    kfx create cluster
    kfx install mla
    kfx install kf
    KUBECONFIG="$("${KIND}" get kubeconfig-path --name "kf-kind")"
    export KUBECONFIG
    kubectl get po -n kubeflow

    # TODO: invoke your tests here
    # teardown will happen automatically on exit
}

main
