package auth

import (
	"context"
	v1 "googelo-shop-gf/app/auth/api/auth/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedAuthServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAuthServiceServer(s.Server, &Controller{})
}

func (*Controller) DeliverTokenByRPC(ctx context.Context, req *v1.DeliverTokenReq) (res *v1.DeliveryResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) VerifyTokenByRPC(ctx context.Context, req *v1.VerifyTokenReq) (res *v1.VerifyResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
