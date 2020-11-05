module github.com/makepeak/steve-operation/payment-web

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/makepeak/steve-operation v0.0.7
	github.com/makepeak/steve-operation/auth v0.0.0-20201103125854-3ede5d4e3464
	github.com/makepeak/steve-operation/payment-service v0.0.0-20201103125854-3ede5d4e3464
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins v1.5.1 // indirect
	github.com/micro/go-plugins/config/source/grpc/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
)
