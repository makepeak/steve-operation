package main

import (
	"fmt"

	"github.com/makepeak/steve-operation/auth/handler"
	"github.com/makepeak/steve-operation/auth/model"
	s "github.com/makepeak/steve-operation/auth/proto/auth"
	"github.com/makepeak/steve-operation/basic"
	"github.com/makepeak/steve-operation/basic/common"
	"github.com/makepeak/steve-operation/basic/config"
	tracer "github.com/makepeak/steve-operation/plugins/tracer/jaeger"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	//"github.com/micro/go-micro/v2/registry/consul"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	openTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"os"
)

var (
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}


func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用 Consul 注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
		micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context)  error{
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
		/*
		// Loads CLI configuration
                micro.Action(func(c *cli.Context) error {
                        return nil
                }),
		*/
	)

	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func initCfg() {
	configAddr := os.Getenv("MICRO_BOOK_CONFIG_GRPC_ADDR")
	source := grpc.NewSource(
		grpc.WithAddress(configAddr),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}


