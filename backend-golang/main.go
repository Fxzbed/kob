package main

import (
	_ "backend-golang/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"backend-golang/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
