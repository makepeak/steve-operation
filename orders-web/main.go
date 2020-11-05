package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/makepeak/steve-operation/basic"
	"github.com/makepeak/steve-operation/basic/common"
	"github.com/makepeak/steve-operation/basic/config"
	"github.com/makepeak/steve-operation/orders-web/handler"
	tracer "github.com/makepeak/steve-operation/plugins/tracer/jaeger"
	"github.com/makepeak/steve-operation/plugins/tracer/opentracing/std2micro"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	//"github.com/micro/go-micro/v2/registry/consul"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	appName = "orders_web"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 创建新服务
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	//设置采样率
	std2micro.SetSamplingFrequency(50)
	// 新建订单接口
	authHandler := http.HandlerFunc(handler.New)
	service.Handle("/orders/new", std2micro.TracerWrapper(handler.AuthWrapper(authHandler)))
	service.Handle("/", std2micro.TracerWrapper(http.HandlerFunc(handler.Hello)))

	// 运行服务
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
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
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

