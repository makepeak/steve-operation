module github.com/makepeak/steve-operation/user-web

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/makepeak/steve-operation v0.0.8
	github.com/makepeak/steve-operation/auth v0.0.0-20201103125854-3ede5d4e3464
	github.com/makepeak/steve-operation/user-service v0.0.0-20201102085203-45070b885e92
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/grpc/v2 v2.0.3
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
)
