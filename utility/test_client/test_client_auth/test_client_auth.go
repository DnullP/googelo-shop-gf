package main

import (
	v1 "googelo-shop-gf/app/auth/api/auth/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/net/gtrace"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()

	serviceName := "auth-service"
	endpoint := "127.0.0.1:4317"
	traceToken := "******_******"

	shutdown, err := otlpgrpc.Init(serviceName, endpoint, traceToken)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	defer shutdown(ctx)

	ctx, span := gtrace.NewSpan(gctx.New(), "StartRequests")
	defer span.End()

	var (
		conn   = grpcx.Client.MustNewGrpcClientConn("demo")
		client = v1.NewAuthServiceClient(conn)
	)

	res, err := client.DeliverTokenByRPC(ctx, &v1.DeliverTokenReq{UserId: 123})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	g.Log().Debug(ctx, "Response:", res.Token)

}
