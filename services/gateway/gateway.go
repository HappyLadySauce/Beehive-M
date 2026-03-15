// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/HappyLadySauce/Beehive-M/pkg/code"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/handler"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

var configFile = flag.String("f", "etc/gateway-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 非 CodeMsg 错误统一为 code/msg 格式
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		if e, ok := err.(*errors.CodeMsg); ok {
			return http.StatusOK, map[string]interface{}{"code": e.Code, "msg": e.Msg}
		}
		return http.StatusOK, map[string]interface{}{"code": code.CodeInternal, "msg": err.Error()}
	})

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
