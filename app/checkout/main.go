package main

import (
	_ "googelo-shop-gf/app/checkout/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"googelo-shop-gf/app/checkout/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
