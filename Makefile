
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 \
	       --proto_path=. \
	       --micro_out=${MODIFY}:. \
	       --go_out=${MODIFY}:. \
	       --validate_out=lang=go:. \
	       proto/lost_found/lost_found.proto

.PHONY: build
build:

	go build -o lost-found-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t lost-found-service:latest
