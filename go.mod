module github.com/XMUMY/lost_found

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/XMUMY/api v0.0.2-0.20200910093640-a2fb08c315d1
	github.com/XMUMY/lib v0.0.0-20200902103721-88651601333c
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.4.1
)
