ARG GOLANG_VERSION=1.12.7
FROM golang:$GOLANG_VERSION as builder

RUN apt-get update
RUN apt-get install -y git unzip

ENV PATH /go/bin:/usr/local/go/bin:${PATH}

ENV GO111MODULE=on
ENV GOPATH=/go

RUN mkdir -p ${GOPATH}/src/github.com/CiscoAI/create-kf-app
WORKDIR ${GOPATH}/src/github.com/CiscoAI/create-kf-app

# Download dependencies first to optimize Docker caching.
COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy in the source
COPY . .

FROM builder as create-kf-app

RUN make build

FROM alpine:3.10.1 as barebones_base
RUN mkdir -p /opt/kubeflow
WORKDIR /opt/kubeflow

FROM barebones_base as kfctl

COPY --from=create-kf-app /go/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-linux /usr/local/bin

CMD ["/bin/bash", "-c", "trap : TERM INT; sleep infinity & wait"]
