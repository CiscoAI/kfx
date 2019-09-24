GOPATH ?= $(HOME)/go
TAG ?= 0.1
OPSYS ?= darwin

build: go-install build-create-kf-app build-kfx

go-install:
	go mod vendor
	go mod download
	go get -u github.com/go-bindata/go-bindata/...
	go-bindata -o pkg/manifests/manifests.go -pkg=manifests manifests
build-create-kf-app:
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-linux ${GOPATH}/src/github.com/CiscoAI/create-kf-app/cli/create-kf-app.go
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-darwin ${GOPATH}/src/github.com/CiscoAI/create-kf-app/cli/create-kf-app.go
build-kfx:
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-linux ${GOPATH}/src/github.com/CiscoAI/create-kf-app/kfx/kfx.go
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-darwin ${GOPATH}/src/github.com/CiscoAI/create-kf-app/kfx/kfx.go	
test: 
	${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-${OPSYS} init --name kf-test
	${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-${OPSYS} delete cluster --name kf-test		
clean:
	packr2 clean
tar:
	tar -vzcf ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-linux.tar.gz ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-linux
	tar -vzcf ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-darwin.tar.gz ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-darwin
	tar -zvcf ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-linux.tar.gz ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-linux
	tar -zvcf ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-darwin.tar.gz ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-darwin
release: tar
	