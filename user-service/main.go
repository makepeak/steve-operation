package main

import (
	"fmt"

	"github.com/makepeak/steve-operation/user-service/basic"
	"github.com/makepeak/steve-operation/user-service/basic/config"
	"github.com/makepeak/steve-operation/user-service/handler"
	"github.com/makepeak/steve-operation/user-service/model"
	s "github.com/makepeak/steve-operation/user-service/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	// micReg := etcd.NewRegistry(registryOptions)

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)
	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.service.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// registryOptions consul获取配置的ip port信息
func registryOptions(ops *registry.Options) {
	// etcdCfg := config.GetEtcdConfig()
	consulCfg := config.GetConsulConfig() 
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}

