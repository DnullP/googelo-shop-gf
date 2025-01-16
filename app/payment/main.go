package main

import (
	_ "googelo-shop-gf/app/payment/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"googelo-shop-gf/app/payment/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
