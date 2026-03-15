package main

import (
	"flag"
	"fmt"

	"github.com/HappyLadySauce/Beehive-M/services/presence/internal/config"
	presenceserviceServer "github.com/HappyLadySauce/Beehive-M/services/presence/internal/server/presenceservice"
	"github.com/HappyLadySauce/Beehive-M/services/presence/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/presence/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/presence.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPresenceServiceServer(grpcServer, presenceserviceServer.NewPresenceServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
