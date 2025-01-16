package main

import (
	_ "googelo-shop-gf/app/product/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"googelo-shop-gf/app/product/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
