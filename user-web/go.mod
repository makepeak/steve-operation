module github.com/makepeak/steve-operation/user-web

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/makepeak/steve-operation/user-service v0.0.0-20201102085203-45070b885e92
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20201009111623-6942a78f3f99
	github.com/micro/go-micro/v2 v2.9.1
)
