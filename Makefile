GOPATH ?= $(HOME)/go
TAG ?= 0.1

build-create-kf-app:
	env GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-linux ${GOPATH}/src/github.com/CiscoAI/create-kf-app/kfx/kfx.go
	env GOOS=darwin GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/create-kf-app-macos ${GOPATH}/src/github.com/CiscoAI/create-kf-app/kfx/kfx.go
build-kfx:
	env GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-linux ${GOPATH}/src/github.com/CiscoAI/create-kf-app/kfx/kfx.go
	env GOOS=darwin GOARCH=amd64 go build -o ${GOPATH}/src/github.com/CiscoAI/create-kf-app/bin/kfx-macos ${GOPATH}/src/github.com/CiscoAI/create-kf-app/kfx/kfx.go	
test:
	ls