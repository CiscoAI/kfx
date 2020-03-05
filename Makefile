GOPATH ?= $(HOME)/go
TAG ?= 0.1
OPSYS ?= darwin

DIR ?= $(shell pwd)
LYFT_IMAGE ?= "lyft/protocgenerator:d53ce1490e235bf765c93b4a8cfcdd07a1325024"

build: go-install build-kfx

gen-proto:
	rm -rf ${DIR}/gen
	docker run -u $(id -u):$(id -g) -v ${DIR}:/defs ${LYFT_IMAGE} -i ./api -d ./api -l go --go_source_relative

go-install:
	go mod vendor
	go mod download
	go get -u github.com/go-bindata/go-bindata/...
	${GOPATH}/bin/go-bindata -o pkg/manifests/manifests.go -pkg=manifests manifests

build-kfx:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/kfx/bin/kfx-linux ${GOPATH}/src/github.com/CiscoAI/kfx/cmd/kfx/kfx.go
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/kfx/bin/kfx-darwin ${GOPATH}/src/github.com/CiscoAI/kfx/cmd/kfx/kfx.go	

test: 
	${GOPATH}/src/github.com/CiscoAI/kfx/bin/create-kf-app-${OPSYS} init --name kf-test
	${GOPATH}/src/github.com/CiscoAI/kfx/bin/create-kf-app-${OPSYS} delete cluster --name kf-test		

tar:
	tar -zvcf ${GOPATH}/src/github.com/CiscoAI/kfx/bin/kfx-linux.tar.gz ${GOPATH}/src/github.com/CiscoAI/kfx/bin/kfx-linux
	tar -zvcf ${GOPATH}/src/github.com/CiscoAI/kfx/bin/kfx-darwin.tar.gz ${GOPATH}/src/github.com/CiscoAI/kfx/bin/kfx-darwin
