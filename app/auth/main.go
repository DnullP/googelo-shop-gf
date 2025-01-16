package main

import (
	"googelo-shop-gf/app/auth/internal/controller/auth"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.New()
	serviceName := "auth-service"
	endpoint, err := g.Cfg().Get(ctx, "Jaeger.address", "127.0.0.1:4317")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	g.Log().Debugf(ctx, "Jaeger at %s", endpoint.String())
	traceToken := "******_******"

	shutdown, err := otlpgrpc.Init(serviceName, endpoint.String(), traceToken)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	defer shutdown(ctx)

	s := grpcx.Server.New()
	auth.Register(s)
	s.Run()
}
