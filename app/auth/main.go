package main

import (
	"fmt"

	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.New()
	test, err := g.Cfg().Get(ctx, "test", "no")
	if err != nil {

	}
	fmt.Println(test)

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

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write(test)
	})
	s.Run()
}
