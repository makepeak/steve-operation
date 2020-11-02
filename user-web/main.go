package main

import (
        log "github.com/micro/go-micro/v2/logger"
	"net/http"
        "github.com/micro/go-micro/v2/web"
        "github.com/makepeak/steve-operation/user-web/handler"

	"github.com/micro/go-micro/v2/registry"
        "github.com/micro/go-plugins/registry/consul/v2"
)

var consulReg registry.Registry

func init(){
    //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
    consulReg = consul.NewRegistry(
        registry.Addrs("127.0.0.1:8500"),
    )
}


func main() {
	// create new web service
        service := web.NewService(
                web.Name("steve.user.web.user"),
                web.Version("latest"),
		web.Address("0.0.0.0:9090"),
		web.Registry(consulReg),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/user/call", handler.Login)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
