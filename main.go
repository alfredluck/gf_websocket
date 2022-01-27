package main

import (
	"gf_websocket/internal/cmd"
	_ "gf_websocket/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {

	cmd.Main.Run(gctx.New())

}
