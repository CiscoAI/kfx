ARG GOLANG_VERSION=1.13
FROM golang:$GOLANG_VERSION as builder

ENV PATH /go/bin:/usr/local/go/bin:${PATH}

ENV GO111MODULE=on
ENV GOPATH=/go

RUN mkdir -p ${GOPATH}/src/github.com/CiscoAI/kfx
WORKDIR ${GOPATH}/src/github.com/CiscoAI/kfx

# Download dependencies first to optimize Docker caching.
COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy in the source
COPY . .

RUN	go get -u github.com/go-bindata/go-bindata/...
RUN	${GOPATH}/bin/go-bindata -o pkg/manifests/manifests.go -pkg=manifests manifests

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kfx -ldflags "-w" -a cmd/kfx/kfx.go

FROM scratch

WORKDIR /

COPY --from=builder /go/src/github.com/CiscoAI/kfx .

ENTRYPOINT ["/kfx"]
